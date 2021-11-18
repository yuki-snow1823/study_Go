package main

import (
	"fmt"
	"strconv"
	"github.com/yuki-snow1823/go_for_paiza" // 外部から呼び出している
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

  r := 0

	for ; r < 5; {
		r++
		fmt.Println("hoge")
	}

	wordset := [...]string{"apple", "banana", "grape"}

	for _, word := range wordset {
		fmt.Printf(word)
	}
	
	str := go_for_paiza.Hello()
	fmt.Print(str)
}
