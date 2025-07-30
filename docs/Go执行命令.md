# Go
### Go执行命令
```bash
go mod init .
go mod init github.com/your_username/your_project_name
go mod init your_project_name
go mod tidy
go run main.go
go build -o your_executable_name
go build -o your_executable_name main.go
go install
go get github.com/your_username/your_package_name
go get -u github.com/your_username/your_package_name
go get -u all
go get -u github.com/your_username/your_package_name@latest

go env -w GOPROXY=https://goproxy.cn,direct
```