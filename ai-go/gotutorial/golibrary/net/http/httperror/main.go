/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:51:49
*/
package main

import (
	"net/http"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func main() {

}
