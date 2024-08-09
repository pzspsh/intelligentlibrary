/*
@File   : pocexec.go
@Author : pan
@Time   : 2023-06-12 10:43:48
*/
package pocexec

type PocOptions struct {
	Proxy      string
	Target     string
	ScriptType string
	Host       string
	Port       int
}

func (p *PocOptions) PocExec() {

}
