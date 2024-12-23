/*
@File   : output.go
@Author : pan
@Time   : 2023-06-12 10:38:37
*/
package output

type Output struct {
	ID         string
	Name       string
	ScriptPath string
	Request    string
	Headers    map[string]string
	Response   string
	Body       string
}
