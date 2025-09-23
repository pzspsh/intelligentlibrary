package main

import (
	"net/http"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func main() {

}
