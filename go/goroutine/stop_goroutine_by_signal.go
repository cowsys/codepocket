package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	done := make(chan struct{})
	quit := make(chan struct{})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)

	go doLoop(quit, done)

	go func() {
		// signalを待つ
		<-sig
		quit <- struct{}{}
	}()

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
