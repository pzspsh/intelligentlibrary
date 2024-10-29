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