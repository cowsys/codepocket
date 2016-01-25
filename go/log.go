package main

import "log"
import "fmt"

func main() {
	fmt.Println("start")

	log.Print("testtest")
	log.Println("testtest")

	// set Prefix
	log.SetPrefix("prefix")
	log.Print("testtest")
	log.Println("testtest")
	log.SetPrefix("")
	log.Print("testtest")
	log.Println("testtest")

	// fatal
	// log.Fatal("testtest")
	// panic
	// log.Panic("testtest")

	fmt.Println("end")
}
