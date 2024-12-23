/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:11:47
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type InstanceDetail struct {
	Entries []InstanceEntries `json:"entries"`
}

type InstanceEntries struct {
	Values InstanceEntriesValues `json:"values"`
}

type InstanceEntriesValues struct {
	IncidentNumber      string `json:"Incident Number"`
	Description         string `json:"Description"`
	ExternalReason      string `json:"ExternalReason"`
	DetailedDescription string `json:"Detailed Description"`
	ExternalName        string `json:"ExternalName"`
}

/*
结构体对应带数组的json
// 队列示例服务server/main.go
*/
func main() {
	url := "http://127.0.0.1:8080/post"
	s := InstanceDetail{Entries: []InstanceEntries{{Values: InstanceEntriesValues{
		IncidentNumber:      "0000000001",
		Description:         "Test",
		DetailedDescription: "Test",
		ExternalName:        "ExternalName",
		ExternalReason:      "EscaltionReason",
	}},
		{InstanceEntriesValues{
			IncidentNumber:      "0000000002",
			Description:         "Test",
			DetailedDescription: "Test",
			ExternalName:        "ExternalName",
			ExternalReason:      "EscaltionReason",
		}}}}

	js, err := json.MarshalIndent(s, "", "   ") // 将json形式的字符串进行格式化
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
