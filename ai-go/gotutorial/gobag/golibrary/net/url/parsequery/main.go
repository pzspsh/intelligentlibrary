/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:30:16
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	m, err := url.ParseQuery(`x=1&y=2&y=3`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toJSON(m))
}

func toJSON(m any) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}
