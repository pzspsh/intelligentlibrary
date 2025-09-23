package main

import (
	"fmt"
	"net/http"
)

func main() {
	targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"
	req, _ := http.NewRequest("GET", targetUrl, nil)
	req.Header.Add("Authorization", "xxxx")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
}
