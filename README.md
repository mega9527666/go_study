
go mod init mega

go get github.com/go-sql-driver/mysql


go build -o mega_go_webserver webserver/webserver.go

go run ./test/test.go 


./build/go_webserver 9600 1