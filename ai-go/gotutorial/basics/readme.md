# 基础(basics)

```go
func ReadFile(filepath string) (data []byte, err error) {
	data, err = os.ReadFile(filepath)
	if err != nil {
		return
	}
	return
}
```

```go
for range 10 {
	fmt.Println("hello")
}
```

```go
type ServiceOptions struct {
	Target  Target
	Proxy   int
	Rate    int
	Timeout int
	Threads int
}

type Target struct {
	Host     string
	Port     int
	Protocol string
}

func (s ServiceOptions) ServiceScan(host string, port int) output.AssetInfo {
	s.Target.Host = host
	s.Target.Port = port
	s.Target.Protocol = "tcp"
	// 修改r.Target的值都会成功，因为传递的是值拷贝，而不是指针拷贝。
}

func (s *ServiceOptions) ServiceScan(host string, port int) output.AssetInfo {
	s.Target.Host = host
	s.Target.Port = port
	s.Target.Protocol = "tcp"
	// 修改r.Target的值不一定成功，因为传递的是指针拷贝，而不是值拷贝。
}
```