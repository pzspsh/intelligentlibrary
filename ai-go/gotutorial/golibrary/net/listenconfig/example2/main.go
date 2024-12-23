/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:06:57
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	upgrade bool
	ln      net.Listener
	server  *http.Server
)

func init() {
	flag.BoolVar(&upgrade, "upgrade", false, "user can't use this")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from pid:%d, ppid: %d\n", os.Getpid(), os.Getppid())
}

func setupSignal() {
	ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	switch sig {
	// case syscall.SIGUSR2:
	// 	err := forkProcess()
	// 	if err != nil {
	// 		fmt.Printf("fork process error: %s\n", err)
	// 	}
	// 	err = server.Shutdown(context.Background())
	// 	if err != nil {
	// 		fmt.Printf("shutdown after forking process error: %s\n", err)
	// 	}
	case syscall.SIGINT, syscall.SIGTERM:
		signal.Stop(ch)
		close(ch)
		err := server.Shutdown(context.Background())
		if err != nil {
			fmt.Printf("shutdown error: %s\n", err)
		}
	}
}

func ForkProcess() error {
	flags := []string{"-upgrade"}
	cmd := exec.Command(os.Args[0], flags...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	l, _ := ln.(*net.TCPListener)
	lfd, err := l.File()
	if err != nil {
		return err
	}
	cmd.ExtraFiles = []*os.File{lfd}
	return cmd.Start()
}

func main() {
	flag.Parse()
	http.HandleFunc("/", hello)
	server = &http.Server{Addr: ":8999"}
	var err error
	if upgrade {
		fd := os.NewFile(3, "")
		ln, err = net.FileListener(fd)
		if err != nil {
			fmt.Printf("fileListener fail, error: %s\n", err)
			os.Exit(1)
		}
		fd.Close()
	} else {
		ln, err = net.Listen("tcp", server.Addr)
		if err != nil {
			fmt.Printf("listen %s fail, error: %s\n", server.Addr, err)
			os.Exit(1)
		}
	}
	go func() {
		err := server.Serve(ln)
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("serve error: %s\n", err)
		}
	}()
	setupSignal()
	fmt.Println("over")
}
