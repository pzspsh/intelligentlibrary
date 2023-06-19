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
