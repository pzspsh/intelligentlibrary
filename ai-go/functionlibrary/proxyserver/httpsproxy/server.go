/*
@File   : server2.go
@Author : pan
@Time   : 2023-11-09 15:09:31
*/
package httpsproxy

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stderr, "httpsproxy:", log.Llongfile|log.LstdFlags)

func StartServe(listenAdress string) {
	cert, err := genCertificate()
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Addr:      listenAdress,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Serve(w, r)
		}),
	}

	logger.Fatal(server.ListenAndServe())

}

func genCertificate() (cert tls.Certificate, err error) {
	rawCert, rawKey, err := generateKeyPair()
	if err != nil {
		return
	}
	return tls.X509KeyPair(rawCert, rawKey)

}

func generateKeyPair() (rawCert, rawKey []byte, err error) {
	// Create private key and self-signed certificate
	// Adapted from https://golang.org/src/crypto/tls/generate_cert.go

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	validFor := time.Hour * 24 * 365 * 10 // ten years
	notBefore := time.Now()
	notAfter := notBefore.Add(validFor)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		fmt.Println("rand Int Error:", err)
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Zarten"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return
	}

	rawCert = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	rawKey = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return
}

func Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		handleHttps(w, r)
	} else {
		handleHttp(w, r)
	}

}

func handleHttps(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", r.Host, 60*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)

}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func checkAdress(adress string) bool {
	_, err := net.ResolveTCPAddr("tcp", adress)
	if err != nil {
		fmt.Println("checkAddree error:", err)
		return false
	}
	return true
}

func SartHttps(proxyip, proxyport string) {
	listenAdress := fmt.Sprintf("%s:%s", proxyip, proxyport)
	if !checkAdress(listenAdress) {
		logger.Fatal("-L listen address format incorrect.Please check it")
	}
	StartServe(listenAdress)
}
