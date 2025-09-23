package main

import (
	"fmt"
	"net/http"
)

func main() {
	targetUrl := "https://ddbc5ffb-c596-4f78-a99d-a6ea93bdc14f.mock.pstmn.io/user/1"
	req, _ := http.NewRequest("DELETE", targetUrl, nil)
	req.Header.Add("Authorization", "xxxx")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
}
