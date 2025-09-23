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
