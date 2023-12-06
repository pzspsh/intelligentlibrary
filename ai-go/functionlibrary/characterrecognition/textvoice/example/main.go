/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 14:53:04
*/
package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("A.wav") //打开要识别的语音文件
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, b)
	if err != nil {
		panic(err)
	}

	//   rate := []string{"8000", "11025", "16000", "22050" ,"24000", "32000", "44100","48000"}google识别的音频频率

	body, err := getText(buf, "22050") //我录制的频率是22050
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func getText(bodys io.Reader, rate string) ([]byte, error) {
	req, err := http.NewRequest("POST", "http://www.google.com/speech-api/v1/recognize?xjerr=1&client=chromium&maxresults=1&lang=zh-CN", bodys)
	if err != nil {
		return nil, err
	}
	var httpclient *http.Client = &http.Client{}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "audio/L16;rate="+rate)
	resp, err := httpclient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body

	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("handlePost statuscode=%d, body=%s", resp.StatusCode, body)
		return nil, errors.New(msg)
	}
	return body, nil
}
