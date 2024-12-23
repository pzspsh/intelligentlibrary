/*
@File   : templates.go
@Author : pan
@Time   : 2023-06-12 09:57:09
*/
package templates

type Templates struct {
	ID            string `yaml:"id" json:"id" jsonschema:"title=id"`
	Info          string `yaml:"info,omitempty" json:"info,omitempty"`
	HttpRequest   string `yaml:"http,omitempty" json:"http,omitempty"`
	SelfContained bool
	Variabless map[string]string
	
}
