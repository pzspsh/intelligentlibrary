/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 14:42:16
*/
package main

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// 文本转语音
func main() {
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("SAPI.SpVoice")
	voice, _ := unknown.QueryInterface(ole.IID_IDispatch)
	saveFile, _ := oleutil.CreateObject("SAPI.SpFileStream")
	ff, _ := saveFile.QueryInterface(ole.IID_IDispatch)
	// 打开wav文件
	oleutil.CallMethod(ff, "Open", "E:\\mygo\\aa.wav", 3, true)
	// 设置voice的AudioOutputStream属性，必须是PutPropertyRef，如果是PutProperty就无法生效
	oleutil.PutPropertyRef(voice, "AudioOutputStream", ff)
	// 设置语速
	oleutil.PutProperty(voice, "Rate", -3)
	// 设置音量
	oleutil.PutProperty(voice, "Volume", 200)
	// 说话
	oleutil.CallMethod(voice, "Speak", "您有新工单，请及时处理！")
	oleutil.CallMethod(voice, "Speak", "bb", 1)
	// 停止说话
	//oleutil.CallMethod(voice, "Pause")
	// 恢复说话
	//oleutil.CallMethod(voice, "Resume")
	// 等待结束
	oleutil.CallMethod(voice, "WaitUntilDone", 1000000)
	// 关闭文件
	oleutil.CallMethod(ff, "Close")
	ff.Release()
	voice.Release()
	ole.CoUninitialize()
}
