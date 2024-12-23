/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:35:34
*/
package main

/*
创建证书

1.创建根证书私钥长度为2048
openssl genrsa -out ca.key 2048

2.利用私钥创建根证书按照提示一路输入：
openssl req -new -x509 -days 36500 -key ca.key -out ca.crt

3.创建长度为2048的SSL证书私匙
openssl genrsa -out server.key 2048

4.利用刚才的私匙建立SSL证书请求一路向下：
openssl req -new -key server.key -out server.csr

5.当前文件夹下运行如下命令创建所需目录：
mkdir dir demoCA &&cd demoCA&&mkdir newcerts&&echo '10' > serial &&touch index.txt&&cd ..

6.用CA根证书签署SSL自建证书：
openssl ca -in server.csr -out server.crt -cert ca.crt -keyfile ca.key

7.查看证书
openssl x509 -in server.crt -noout -text
*/
import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func HandleClientConnect(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Receive Connect Request From ", conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	for {
		len, err := conn.Read(buffer)
		if err != nil {
			log.Println(err.Error())
			break
		}
		fmt.Printf("Receive Data: %s\n", string(buffer[:len]))
		//发送给客户端
		_, err = conn.Write([]byte("服务器收到数据:" + string(buffer[:len])))
		if err != nil {
			break
		}
	}
	fmt.Println("Client " + conn.RemoteAddr().String() + " Connection Closed.....")
}

func main() {
	crt, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalln(err.Error())
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	// Time returns the current time as the number of seconds since the epoch.
	// If Time is nil, TLS uses time.Now.
	tlsConfig.Time = time.Now
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.
	// The Reader must be safe for use by multiple goroutines.
	tlsConfig.Rand = rand.Reader
	l, err := tls.Listen("tcp", ":8888", tlsConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			go HandleClientConnect(conn)
		}
	}

}
