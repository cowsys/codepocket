package main

import "fmt"

type Num int

// 連番の定数を定義する1
const (
	One Num = 1 + iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Eleven
)

// 連番の定数を定義する2
const (
	XOne Num = iota
	XTwo
	XThree
	XFour
	XFive
	XSix
	XSeven
	XEight
	XNine
	XTen
	XEleven
)

// constのalignは有効になる(gofmtに潰されない)
const (
	Piyo     Num = 1
	Piyopiyo Num = 2
)

// iotaを利用した定数定義と, 配列による文字列の定義を使うと
// マッピングをうまく定義できる
var numbers = [...]string{
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
}

func main() {
	// 1+iota(1からスタート)
	{
		fmt.Println(One)
		fmt.Println(Two)
		fmt.Println(Three)
		fmt.Println(Four)
		fmt.Println(Five)
		fmt.Println(Six)
		fmt.Println(Seven)
		fmt.Println(Eight)
		fmt.Println(Nine)
		fmt.Println(Ten)
		fmt.Println(Eleven)
	}
	fmt.Println("----")
	// iota(0からスタート)
	{
		fmt.Println(XOne)
		fmt.Println(XTwo)
		fmt.Println(XThree)
		fmt.Println(XFour)
		fmt.Println(XFive)
		fmt.Println(XSix)
		fmt.Println(XSeven)
		fmt.Println(XEight)
		fmt.Println(XNine)
		fmt.Println(XTen)
		fmt.Println(XEleven)

	}
	fmt.Println("----")
	{
		// iotaのみのときは
		// そのままindexに指定すれば値を取得可能
		fmt.Println(numbers[XOne])
		fmt.Println(numbers[XFive])
		fmt.Println(numbers[XEleven])

		fmt.Println("--")
		// iota+1のときは
		// -1
		fmt.Println(numbers[One-1])
		fmt.Println(numbers[Five-1])
		fmt.Println(numbers[Eleven-1])
	}
}
