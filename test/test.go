package main

import (
	"fmt"
)

const (
	None = iota
	Male
	Femail
)

func main() {
	fmt.Println("abcd")
	// var a string = "abcde"
	// var b int
	// var c bool
	// var d float64
	// var e string
	// fmt.Println("a=", a)
	// fmt.Println("b=", b)
	// fmt.Println("c=", c)
	// fmt.Println("d=", d)
	// fmt.Println("e=", e)
	// ewrere, numb, strs := numbers()
	// fmt.Println("numbers====", ewrere, numb, strs)
	const m, n, o = 1, 2, false
	// fmt.Println("mno=", m, n, o)
	// fmt.Println("Male====", Male, Femail, None)
	// fmt.Println("len==", len(a))
	// fmt.Println("sizeof=", unsafe.Sizeof(a))
	// fmt.Println("max=", max(1, 2))

	var numbers [5]int
	var numbers1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers", numbers[0], numbers)
	fmt.Println("numbers1", numbers1[0], numbers1)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 100
	}
	fmt.Println("numbers===2", numbers[0], numbers)

}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "abcdefg"
	return a, b, c
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
