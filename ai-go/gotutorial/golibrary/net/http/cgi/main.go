package main

import (
	"fmt"
	"net/http"
	"time"
)

func Demo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.go")
	var closenotify http.CloseNotifier
	closenotify = w.(http.CloseNotifier)
	select {
	case <-closenotify.CloseNotify():
		fmt.Println("cut")
	case <-time.After(time.Duration(100) * time.Second):
		fmt.Println("timeout")

	}
}
func main() {
	http.HandleFunc("/test", Demo)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}

}
