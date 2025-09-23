package expexec

type ExpOptions struct {
	Cmd        string
	Proxy      string
	Target     string
	Host       string
	Port       int
	ScriptType string
	Username   string
	Password   string
}

func (e *ExpOptions) ExpExec() {

}
