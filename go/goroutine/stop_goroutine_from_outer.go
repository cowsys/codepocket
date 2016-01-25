package main

import "fmt"
import "time"

func main() {
	//// NOTE:中から外に終了を伝えるためのchannel
	done := make(chan struct{})

	// 外部からループするgoroutineを停止させる
	//// NOTE:外から中に終了を伝えるためのchannel
	quit := make(chan struct{})
	go doLoop(quit, done)
	{
		time.Sleep(time.Second)
		fmt.Println("1 second passed")
		quit <- struct{}{}
	}
	<-done
}

func doLoop(quit, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	loop := true
	go func() {
		defer fmt.Println("doLoop exited.")
		<-quit
		loop = false
	}()

	for loop {
		fmt.Println("loop continued...")
	}
}
