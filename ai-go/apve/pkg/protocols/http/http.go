/*
@File   : http.go
@Author : pan
@Time   : 2023-06-17 17:39:29
*/
package http

type HttpRuquests struct {
	Method             string                 `yaml:"method,omitempty"`
	Path               string                 `yaml:"path,omitempty"`
	Body               string                 `yaml:"body,omitempty"`
	Headers            map[string]string      `yaml:"headers,omitempty"`
	StopAtFirstMatch   bool                   `yaml:"stop-at-first-match,omitempty"`
	MatchersCondition  string                 `yaml:"matchers-condition,omitempty"`
	Raw                []string               `yaml:"raw,omitempty"`
	Attack             string                 `yaml:"attack,omitempty"`
	Payloads           map[string]interface{} `yaml:"attack,omitempty"`
	DigestUsername     string
	DigestPassword     string
	IterateAll         bool
	SkipVariablesCheck bool
	CookieRease        bool
	HostRedirects      bool
	MaxRedirects       bool
	ReqCondition       bool
	Threads            int
	MaxSize            int64
	Unsafe             bool
	ReadAll            bool
	Matchers           string
	Extractors         string
}
