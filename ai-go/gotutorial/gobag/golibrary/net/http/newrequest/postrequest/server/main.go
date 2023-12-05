/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:14:57
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func main() {
	http.HandleFunc("/post", postHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}

	fmt.Println("r.Body:", string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))

	var data InstanceDetail
	json.Unmarshal(b, &data)
	//因为entries是个数组所以用数组(下标)方式访问，对象用点（"."）访问
	fmt.Println(data.Entries[0].Values.IncidentNumber)
	fmt.Println(data.Entries[0].Values.Description)
	fmt.Println(data.Entries[1].Values.IncidentNumber)
	fmt.Println(data.Entries[1].Values.Description)
	//如果这里数组的元素数量不固定时，需要使用for循环去获取数组元素的值
	//for _, v := range data.Entries {
	//	fmt.Println("Incident Number:", v.Values.IncidentNumber)
	//	fmt.Println("Description:", v.Values.Description)
	//}
}
