# header头设置

### header.Set
```go
req.Header.Set("User-Agent","自定义的浏览器")
req.Header.Set("User-Agent","自定义的浏览器3")
```
当我们使用Set时候，如果原来这一项已存在，后面的就修改已有的。所以这里最终的结果就是自定义的浏览器3

### header.Add
```go
req.Header.Add("User-Agent","自定义的浏览器")
req.Header.Add("User-Agent","自定义的浏览器3")
```

```go
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.WriteHeader(http.StatusOK)
```
当使用Add时候，如果原本不存在，则添加，如果已存在，就不做任何修改。所以这里最终的结果就是自定义的浏览器

```go

```