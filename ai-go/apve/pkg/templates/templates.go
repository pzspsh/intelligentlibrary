package templates

type Templates struct {
	ID            string `yaml:"id" json:"id" jsonschema:"title=id"`
	Info          string `yaml:"info,omitempty" json:"info,omitempty"`
	HttpRequest   string `yaml:"http,omitempty" json:"http,omitempty"`
	SelfContained bool
	Variabless    map[string]string
}
