/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:09:08
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(c)

	// Block until any signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
