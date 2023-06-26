package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func recvProc(conn net.Conn, out io.Writer) {
	for {
		var buf [512]byte
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			break
		}
		out.Write(buf[:n])
	}
	os.Exit(0)
}

func pushFile(fn string, rawconn io.ReadWriteCloser) error {
	info, err := os.Stat(fn)
	if err != nil {
		return err
	}
	size := info.Size()
	name := filepath.Base(fn)

	fp, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fp.Close()

	conn := textproto.NewConn(rawconn)
	defer conn.Close()

	conn.PrintfLine("Action:Push")
	conn.PrintfLine("Version:1")
	conn.PrintfLine("FileName:%s", name)
	conn.PrintfLine("Size:%d", size)
	conn.PrintfLine("Data")
	_, err = conn.W.ReadFrom(fp)
	if err != nil {
		return err
	}

	return nil
}

func recvFile(rawconn io.ReadWriteCloser) error {
	var name string
	var size int64
	conn := textproto.NewConn(rawconn)
	defer conn.Close()

	line1, err := conn.ReadLine()
	if err != nil {
		return err
	}
	if strings.Compare(line1, "Action:Push") != 0 {
		return fmt.Errorf("error Header:%v", line1)
	}

	line1, err = conn.ReadLine()
	if err != nil {
		return err
	}
	if strings.Compare(line1, "Version:1") != 0 {
		return fmt.Errorf("error Header:%v", line1)
	}

	line1, err = conn.ReadLine()
	if err != nil {
		return err
	}
	if strings.HasPrefix(line1, "FileName:") {
		name = strings.TrimPrefix(line1, "FileName:")
		fmt.Fprintln(os.Stderr, name)
	} else {
		return fmt.Errorf("error Header:%s", line1)
	}

	line1, err = conn.ReadLine()
	if err != nil {
		return err
	}
	if strings.HasPrefix(line1, "Size:") {
		sz := strings.TrimPrefix(line1, "Size:")
		size, err = strconv.ParseInt(sz, 10, 64)
		if err != nil {
			return fmt.Errorf("%w\nHeader:%s", err, line1)
		}
	} else {
		return fmt.Errorf("error Header:%s", line1)
	}

	line1, err = conn.ReadLine()
	if err != nil {
		return err
	}
	if strings.Compare(line1, "Data") != 0 {
		return fmt.Errorf("error Header:%s", line1)
	}

	fp, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fp.Close()
	for size > 0 {
		var buf [512]byte
		n, _ := conn.R.Read(buf[:])
		if n == 0 {
			break
		}
		size -= int64(n)
		if size >= 0 {
			fp.Write(buf[:n])
		} else {
			fp.Write(buf[:n+int(size)])
		}
	}
	return nil
}

// go run netcat.go -p 6666
func main() {
	var port = flag.Int("p", 0, "监听端口")
	var recv = flag.Bool("r", false, "等待接受文件，优先于 -f 参数")
	var fn = flag.String("f", "", "发送文件，如果同时存在 -r 参数，本参数无效")
	flag.Parse()
	var conn net.Conn
	var err error
	if *port == 0 {
		a := flag.Arg(0)
		p := flag.Arg(1)
		conn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", a, p))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
	} else {
		var l net.Listener
		l, err = net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			panic(err)
		}
		defer l.Close()
		showAddr()
		conn, err = l.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		fmt.Fprintln(os.Stderr, "Connected:", conn.RemoteAddr().String())
	}

	if *recv {
		err = recvFile(conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error:%v", err.Error())
		}
		return
	}

	if *fn != "" {
		err = pushFile(*fn, conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error:%v", err.Error())
		}
		return
	}

	go recvProc(conn, os.Stdout)
	for {
		var buf [4096]byte
		n, err := os.Stdin.Read(buf[:])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		conn.Write(buf[:n])
	}
}

func showAddr() {
	ifs, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, if1 := range ifs {
		addrs, err := if1.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			if strings.HasPrefix(addr.String(), "127.") {
				continue
			} else if strings.Contains(addr.String(), ":") {
				continue
			} else {
				vs := strings.Split(addr.String(), "/")
				fmt.Fprintf(os.Stderr, "IP : %s\n", vs[0])
			}
		}
	}

}
