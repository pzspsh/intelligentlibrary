/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 17:38:47
*/
package main

import (
	"flag"
	"net"
	"os"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func main() {
	help := flag.Bool("help", false, "print usage")
	bind := flag.String("bind", "127.0.0.1:6000", "The address to bind to")
	backend := flag.String("backend", "", "The backend server address")
	flag.Parse()

	logger.Level(zerolog.DebugLevel)

	if *help {
		flag.Usage()
		return
	}

	if *backend == "" {
		flag.Usage()
		return
	}

	if *bind == "" {
		//use default bind
		logger.Info().Str("bind", *bind).Msg("use default bind")
	}

	success, err := RunProxy(*bind, *backend)
	if !success {
		logger.Error().Err(err).Send()
		os.Exit(1)
	}
}

func RunProxy(bind, backend string) (bool, error) {
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		return false, err
	}
	defer listener.Close()
	logger.Info().Str("bind", bind).Str("backend", backend).Msg("tcp-proxy started.")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error().Err(err).Send()
		} else {
			go ConnectionHandler(conn, backend)
		}
	}
}

func ConnectionHandler(conn net.Conn, backend string) {
	logger.Info().Str("conn", conn.RemoteAddr().String()).Msg("client connected.")
	target, err := net.Dial("tcp", backend)
	defer conn.Close()
	if err != nil {
		logger.Error().Err(err).Send()
	} else {
		defer target.Close()
		logger.Info().Str("conn", conn.RemoteAddr().String()).Str("backend", target.LocalAddr().String()).Msg("backend connected.")
		closed := make(chan bool, 2)
		go Proxy(conn, target, closed)
		go Proxy(target, conn, closed)
		<-closed
		logger.Info().Str("conn", conn.RemoteAddr().String()).Msg("Connection closed.")
	}
}

func Proxy(from net.Conn, to net.Conn, closed chan bool) {
	buffer := make([]byte, 4096)
	for {
		n1, err := from.Read(buffer)
		if err != nil {
			closed <- true
			return
		}
		n2, err := to.Write(buffer[:n1])
		logger.Debug().Str("from", from.RemoteAddr().String()).Int("recv", n1).Str("to", to.RemoteAddr().String()).Int("send", n2).Send()
		if err != nil {
			closed <- true
			return
		}
	}
}

/*
使用方式
代理监听9000端口，代理后端服务的8000端口

go run main.go --bind 0.0.0.0:9000 --backend 127.0.0.1:8000


这个简单的小程序主要由三个函数构成
1.RunProxy 启动代理服务，监听bind参数指定的端口，接收客户端请求
2.ConnectionHandler 客户端请求处理，连接backend服务
3.Proxy 数据传输代理，将客户端数据发送到backend服务，将backend数据发送给客户端
*/
