package main

import "fmt"

func main() {
	//      この指定方法だと
	//      定義時に要素名を人力設定しなくて良い
	arr := [...]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
	}
	for i, v := range arr {
		fmt.Println("index", i, ":", v)
	}
}
