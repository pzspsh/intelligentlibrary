/*
@File   : models.go
@Author : pan
@Time   : 2023-06-12 11:03:09
*/
package models

type Info struct {
	Name           string `yaml:"name,omitempty" json:"name,omitempty"`
	Author         string
	Description    string
	Reference      []string
	Classification Classification
	Metadata       Metadata
	Remediation    string
	Tags           string
}

type Classification struct {
	CvssMetrics string
	CvssScore   float64
	CveId       string
	CweId       string
	Cpe         string
	EpssScore   float64
}

type Metadata struct {
	MaxRequest  int64
	Verified    bool
	ShodanQuery string
	GoogleQuery string
	FofaQuery   string
}

type UserInfo struct {
	Username               string // 用户名
	Password               string // 密码
	AssetResultPushTopic   string // 资产探测结果topic队列
	AssetScanReceiveTopic  string // 资产探测任务监听Topic队列
	FingerResultPushTopic  string // 指纹识别结果topic队列
	FingerScanReceiveTopic string // 指纹识别任务监听Topic队列
	VulResultPushTopic     string // 漏洞扫描结果topic队列
	VulScanReceiveTopic    string // 漏洞扫描任务监听Topic队列
	Verify                 bool
}
