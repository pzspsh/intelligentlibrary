/*
@File   : scheduler.go
@Author : pan
@Time   : 2023-06-09 18:14:33
*/
package scheduler

type Options struct {
	ProxyBool bool
	ProxyUrl  string
	Username  []string
	Password  []string
	Target    string
}

type AssetScan struct {
	Flag     int    `json:"flag,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	ScanType string `json:"scantype,omitempty"`
}

type DomainScan struct {
	Flag     int    `json:"flag,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	ScanType string `json:"scantype,omitempty"`
}

type PortScan struct { // 端口探活
	Flag     int    `json:"flag,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	ScanType string `json:"scantype,omitempty"` // 探活扫描类型：全端口、top100、top1000、自定义
}

type FingerScan struct {
	Flag  int    `json:"flag,omitempty"`
	Proxy string `json:"proxy,omitempty"`
}

type ServerScan struct {
	Flag  int    `json:"flag,omitempty"`
	Proxy string `json:"proxy,omitempty"`
}

type VulScan struct {
	Flag     int    `json:"flag,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	ScanType string `json:"scantype,omitempty"` // 漏洞扫描类型：全扫描、自定义、根据指纹识别的指纹进行匹配脚本漏洞扫描、。。。
}

type ExpScan struct {
	Flag  int    `json:"flag,omitempty"`
	Proxy string `json:"proxy,omitempty"`
}

type CrackedScan struct {
	Flag     int    `json:"flag,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	Username []string
	Password []string
}

func Scheduler() {

}
