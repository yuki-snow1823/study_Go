package main

import (
	"fmt"
	"strconv"
)

func main() {
	// UTFの文字コードで変換されているから、任意のものは出ない
	var number int = 1234
	var numberStr string = string(number)
	fmt.Println(numberStr)

	var i int
	var s string = "123"
	i, _ = strconv.Atoi(s)
	fmt.Println(i) // 123が出る

	num := 0
	for n := 0; n < 100000000; n++ {
		num += n
	}
	fmt.Println(num)

	const (
		X int = 1
		Y
		Z // 1が出る
	)

	fmt.Print(Z)
}
