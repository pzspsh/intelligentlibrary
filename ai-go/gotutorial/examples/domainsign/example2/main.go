/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 11:18:42
*/
package main

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
)

func main() {
	engine := gin.Default()
	engine.Use(TlsHandler)
	go routine1()
	go routine()
	select {} // 阻塞程序
}

func TlsHandler(c *gin.Context) {
	logrus.Infoln("Used TLS Handler middleware.")
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		// SSLHost: "", //该字段不配置时，默认由http转发至https同域名下
	})
	err := secureMiddleware.Process(c.Writer, c.Request)
	// If there was an error, do not continue.
	if err != nil {
		logrus.Errorf("There is an error in TLS handle. Error: %s\n", err)
		return
	}
	logrus.Infof("TLS Handler middleware used success. SSLHost is: %s\n", "config.ServerSSLHost")
	c.Next()
}

func routine() {
	engine := gin.Default()
	err := engine.Run(":80")
	if err != nil {
		logrus.Errorf("INIT GIN APP RunListener ERROR: %s\n", err)
		panic(err)
	}
}

func routine1() {
	engine := gin.Default()
	err := engine.RunTLS(":443", "certFile", "keyFile")
	if err != nil {
		logrus.Errorf("INIT GIN APP RunListener ERROR: %s\n", err)
		panic(err)
	}
}

func Handler(c *gin.Context) {
	engine := gin.Default()
	srv := &http.Server{
		Addr:    "config.ServerHost" + ":" + "config.ServerPort",
		Handler: engine.Handler(),
		TLSConfig: &tls.Config{
			NameToCertificate: make(map[string]*tls.Certificate, 0),
		},
	}
	// 第一个域名：example.com
	exampleCert, err := tls.LoadX509KeyPair("httpsCerts/example.com.cert", "httpsCerts/example.com.key")
	if err != nil {
		logrus.Errorf("Cannot find example.com cert & key file. Error is: %s\n", err)
		panic(err)
	}
	srv.TLSConfig.NameToCertificate["example.com"] = &exampleCert
	srv.TLSConfig.NameToCertificate["www.example.com"] = &exampleCert
	// 第二个域名：example2.com
	example2Cert, err := tls.LoadX509KeyPair("httpsCerts/example2.com.cert", "httpsCerts/example2.com.key")
	if err != nil {
		logrus.Errorf("Cannot find example.com cert & key file. Error is: %s\n", err)
		panic(err)
	}
	srv.TLSConfig.NameToCertificate["example2.com"] = &example2Cert
	srv.TLSConfig.NameToCertificate["www.example2.com"] = &example2Cert

	srv.TLSConfig.GetCertificate = func(clientInfo *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if x509Cert, ok := srv.TLSConfig.NameToCertificate[clientInfo.ServerName]; ok {
			return x509Cert, nil
		}
		return nil, errors.New("provided servername is invalid 400")
	}
	err = srv.ListenAndServeTLS("", "") // 这里不需要配置certFile和keyFile，因为在前面我们已经配置了自定义的TLSConfig了
	if err != nil {
		logrus.Errorf("INIT GIN APP RunListener ERROR: %s\n", err)
		panic(err)
	}
}
