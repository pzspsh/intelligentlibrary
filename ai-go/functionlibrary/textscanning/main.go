/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 15:13:53
*/
package main

import (
	"fmt"

	"github.com/levigross/grequests"
)

func main() {
	apiKey := "your_api_key"
	secretKey := "your_secret_key"

	imageURL := "https://example.com/image.jpg"

	accessToken, err := getAccessToken(apiKey, secretKey)
	if err != nil {
		fmt.Println("获取百度AI接口访问令牌失败：", err)
		return
	}

	result, err := recognizeText(imageURL, accessToken)
	if err != nil {
		fmt.Println("文字识别失败：", err)
		return
	}

	fmt.Println(result)
}

// 获取百度AI接口访问令牌
func getAccessToken(apiKey string, secretKey string) (string, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", apiKey, secretKey)

	resp, err := grequests.Get(url, nil)
	if err != nil {
		return "", err
	}

	var jsonResp map[string]interface{}
	err = resp.JSON(&jsonResp)
	if err != nil {
		return "", err
	}

	accessToken := jsonResp["access_token"].(string)
	return accessToken, nil
}

// 执行文字识别
func recognizeText(imageURL string, accessToken string) (string, error) {
	url := "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	params := map[string]string{
		"url":          imageURL,
		"access_token": accessToken,
	}

	resp, err := grequests.Post(url, &grequests.RequestOptions{Headers: headers, Data: params})
	if err != nil {
		return "", err
	}

	var jsonResp map[string]interface{}
	err = resp.JSON(&jsonResp)
	if err != nil {
		return "", err
	}

	wordsResults := jsonResp["words_result"].([]interface{})
	result := ""

	for _, word := range wordsResults {
		result += word.(map[string]interface{})["words"].(string) + " "
	}

	return result, nil
}
