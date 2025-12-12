
go mod init mega

go build -o mega_go_webserver webserver/webserver.go

go run ./test/test.go 