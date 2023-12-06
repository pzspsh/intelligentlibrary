/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 14:55:59
*/
package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/cryptix/wav"
	"github.com/nl8590687/asrt-sdk-go/sdk"
)

var (
	winmm         = syscall.MustLoadDLL("winmm.dll")
	mciSendString = winmm.MustFindProc("mciSendStringW")
)

func MCIWorker(lpstrCommand string, lpstrReturnString string, uReturnLength int, hwndCallback int) uintptr {
	lpstrc, err := syscall.UTF16PtrFromString(lpstrCommand)
	if err != nil {
		fmt.Println(err)
	}
	lpstrr, err := syscall.UTF16PtrFromString(lpstrReturnString)
	if err != nil {
		fmt.Println(err)
	}
	i, _, _ := mciSendString.Call(uintptr(unsafe.Pointer(lpstrc)),
		uintptr(unsafe.Pointer(lpstrr)),
		uintptr(uReturnLength), uintptr(hwndCallback))
	return i
}

func main() {
	fmt.Println("winmm.dll Record Audio to .wav file")

	i := MCIWorker("open new type waveaudio alias capture", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code A: ", i)
	}

	i = MCIWorker("set capture  bitspersample  16", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code A1: ", i)
	}

	i = MCIWorker("set capture  channels   1", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code A2: ", i)
	}

	i = MCIWorker("set capture  samplespersec 16000", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code A3: ", i)
	}

	// i = MCIWorker("set capture  bytespersec   32000", "", 0, 0)
	// if i != 0 {
	// 	log.Fatal("Error Code A4: ", i)
	// }
	i = MCIWorker("set capture  alignment   2", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code A5: ", i)
	}

	i = MCIWorker("record capture", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code B: ", i)
	}
	fmt.Println("Listening...")
	time.Sleep(10 * time.Second)

	i = MCIWorker("save capture mic2.wav", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code C: ", i)
	}
	i = MCIWorker("close capture", "", 0, 0)
	if i != 0 {
		log.Fatal("Error Code D: ", i)
	}

	fmt.Println("Audio saved to mic2.wav")
	testInfo, err := os.Stat("mic2.wav")
	checkErr(err)
	testWav, err := os.Open("mic2.wav")
	checkErr(err)

	wavReader, err := wav.NewReader(testWav, testInfo.Size())
	checkErr(err)

	fmt.Println("mic2.wav info")
	fmt.Println(wavReader) //.wav 详细信息
	httpDemo()
}

func httpDemo() {
	// 初始化
	host := "127.0.0.1" //ASRT 服务器地址
	port := "20001"
	protocol := "http"

	sr := sdk.GetSpeechRecognizer(host, port, protocol)
	// ======================================================
	// 识别文件
	filename := "mic2.wav"
	resultFile, err := sr.RecogniteFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	for index, res := range resultFile {
		fmt.Println("Wav文件语音识别结果 ", index, ":", res.Result)
	}

	byteData := sdk.LoadFile(filename)
	wave, err := sdk.DecodeWav(byteData)
	if err != nil {
		fmt.Println(err)
	}
	// ======================================================
	// 识别一段Wave音频序列
	result, err := sr.Recognite(wave.GetRawSamples(), wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语音识别结果：", result.Result)
	// ======================================================
	// 调用声学模型识别一段Wave音频序列
	result, err = sr.RecogniteSpeech(wave.GetRawSamples(), wave.FrameRate, wave.Channels, wave.SampleWidth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语音识别声学模型结果：", result.Result)
	// ======================================================
	// 调用语言模型1
	pinyinResult := []string{}
	for i := 0; i < len(result.Result.([]interface{})); i += 1 {
		pinyinResult = append(pinyinResult, result.Result.([]interface{})[i].(string))
	}

	result, err = sr.RecogniteLanguage(pinyinResult)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("语言模型结果：", result.Result)
	// ======================================================
	// 调用语言模型2
	sequencePinyin := []string{"ni3", "hao3", "a1"}
	result, err = sr.RecogniteLanguage(sequencePinyin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("语言模型结果：", result.Result)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
