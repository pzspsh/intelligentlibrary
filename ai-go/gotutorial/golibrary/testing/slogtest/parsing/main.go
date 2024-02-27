/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:47:30
*/
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"testing/slogtest"
)

func main() {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				panic(err) // In a real test, use t.Fatal.
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		log.Fatal(err)
	}

}
