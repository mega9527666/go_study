package utils

import "fmt"

func init() {
	fmt.Println("HttpUtil.init")
}

func Request(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
