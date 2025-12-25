
go mod init mega

go get github.com/go-sql-driver/mysql


go build -o mega_go_webserver webserver/webserver.go

go run ./test/test.go 


./build/go_webserver 9600 1

<!-- proto使用安装 -->
<!-- 安装go proto生成插件 -->
go get google.golang.org/protobuf@latest


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


brew install protobuf

protoc --version

protoc \
  --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go \
  --go_out=./proto \
  ./proto/*.proto

