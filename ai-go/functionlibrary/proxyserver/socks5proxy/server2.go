/*
@File   : server2.go
@Author : pan
@Time   : 2023-11-09 14:14:22
*/
package socket5proxy

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func process(conn net.Conn) { // Conn 是一个通用的面向流的网络连接。 多个 goroutine 协程可以同时调用 Conn 上的方法
	defer conn.Close()              // defer 会在函数结束后从后往前触发，Close() 手动关闭连接
	reader := bufio.NewReader(conn) // 把输入的连接转换成只读的带缓冲的流
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err) // RemoteAddr() 返回远程网络地址
		return
	}
	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err) // RemoteAddr() 返回远程网络地址
		return
	}
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) { // 用来建立授权，验证身份
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	ver, err := reader.ReadByte() // ReadByte() 读取并返回单个字节，读取到版本号 VER
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	methodSize, err := reader.ReadByte() // ReadByte() 读取并返回单个字节，读取到支持认证的方法数量 NMETHODS
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)   // 创建一个 method 的缓冲区
	_, err = io.ReadFull(reader, method) // ReadFull() 将reader中的字节准确地读取到 method 中
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}

	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00}) // 返回协议版本号 socks5Ver，建立授权的方式 0x00
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

func connect(reader *bufio.Reader, conn net.Conn) (err error) { // 用来建立连接,进行请求
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf) // ReadFull() 将 reader 中的 len(buf) 个字节准确地读取到 buf 中
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}

	/*读取到 VER、CMD、ATYP，并验证其合法性*/
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", ver)
	}
	addr := ""
	switch atyp {
	case atypIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("IPv6: no supported yet")
	default:
		return errors.New("invalid atyp")
	}
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2]) // binary 实现了数字和字节序列之间的简单转换以及 varints 的编码和解码

	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port)) //Dial() 拨号连接到指定网络上的地址,进行 TCP 连接
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close() // defer 会在函数结束后从后往前触发，Close() 手动关闭连接
	log.Println("dial", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0}) // 返回请求报文
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	/*
	   WithCancel() 返回具有新 Done 通道的 parent 副本。返回的上下文的完成通道在调用返回的取消函数或父上下文的完成通道关闭时关闭，以先发生者为准，
	   取消此上下文会释放与其关联的资源，因此代码应在此上下文中运行的操作完成后立即调用取消
	*/
	ctx, cancel := context.WithCancel(context.Background()) // Background() 返回一个非零、空的Context
	defer cancel()

	/* 启动2个协程，实现双向数值转换 */
	/* 从浏览器到服务器 */
	go func() {
		_, _ = io.Copy(dest, reader) // 将副本从 reader 复制到 dest，直到在 reader 上达到EOF或发生错误，它返回复制的字节数和复制时遇到的第一个错误（如果有）
		cancel()
	}()
	/* 从服务器到浏览器 */
	go func() {
		_, _ = io.Copy(conn, dest) // 将副本从 dest 复制到 conn，直到在 dest 上达到EOF或发生错误，它返回复制的字节数和复制时遇到的第一个错误（如果有）
		cancel()
	}()

	<-ctx.Done() // Done() 代表此上下文完成的工作应该被取消时，Done 返回一个关闭的通道
	return nil
}

func StartProxyServer2(proxyip, proxyport string) { //此在SOCKS5 代理程序需要在命令行运行，如输入 go run main.go 和 curl.exe --socks5 127.0.0.1:1080 -v http://www.baidu.com
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%s", proxyip, proxyport)) // Listen()可以收听本地网络地址上的广播，以此来侦听一个端口
	if err != nil {
		panic(err) // panic() 会停止当前 goroutine 协程的正常执行
	}
	for {
		client, err := server.Accept() // Accept() 等待并将下一个连接返回给侦听器
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client) //go 启动一个子线程或 goroutine 协程（处理并发）来处理连接
	}
}
