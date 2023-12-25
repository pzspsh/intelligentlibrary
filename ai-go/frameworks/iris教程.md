# iris使用教程
```go
https://github.com/kataras/iris 
https://iris-go.com
https://github.com/kataras/iris/tree/master/_examples

go get -u github.com/kataras/iris/v12

iris version: 12.1.8

iris run main.go

curl -X POST http://localhost:8080/message -d "username=admin&message=hello world"

curl -X POST http://localhost:8080/message -d "username=admin&message=hello world" -H "Content-Type: application/json"

curl -X POST http://localhost:8080/message -d '{"username":"admin","message":"hello world"}' -H "Content-Type: application/json"
```