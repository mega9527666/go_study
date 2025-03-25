package logger

import "fmt"

func init() {
	fmt.Println("Logger.init")
}

func Log(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
