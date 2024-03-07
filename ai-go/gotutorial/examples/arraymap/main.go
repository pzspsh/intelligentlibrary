/*
@File   : main.go
@Author : pan
@Time   : 2023-10-16 15:31:18
*/
package main

import (
	"encoding/json"
	"fmt"
)

type MapStrings []map[string]string

type ABS struct {
	A1 string `json:"a1,omitempty"`
	A2 string `json:"a2,omitempty"`
}

func main() {
	// bb := []map[string]string{}
	// aa := map[string]string{}
	// cc := map[string]string{}
	// aa["a1"] = "111"
	// aa["a2"] = "222"
	// cc["c1"] = "333"
	// cc["c2"] = "444"
	// bb = append(bb, aa)
	// bb = append(bb, cc)
	// fmt.Println(bb[0]["a2"])

	// astr := `[{"a1":"111", "a2":"222"}, {"a1": "333", "a2": "444"}]`
	// m := []map[string]string{}
	// if err := json.Unmarshal([]byte(astr), &m); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(m)
	// }

	astr := `[{"a1":"111", "a2":"222"}, {"a1": "333", "a2": "444"}]`
	abs := []ABS{}
	if err := json.Unmarshal([]byte(astr), &abs); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(abs)
	}
}
