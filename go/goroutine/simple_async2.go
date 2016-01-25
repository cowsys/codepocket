package main

import "fmt"
import "time"

func main() {
	defer fmt.Println("main exists")

	done := make(chan struct{})
	go func() {
		defer fmt.Println("deferred function fired!")
		// NOTE:mainで定義された
		// goroutine終了判定用channelを通して
		// goroutineの終了をメイン関数に伝える＆
		// メイン関数はそれを待つことで処理同期を実現
		defer func() { done <- struct{}{} }()

		fmt.Println("goroutine starts...")
		time.Sleep(time.Second)
	}()
	<-done

	fmt.Println("sleeping for 0.5 secs...")
	time.Sleep(500 * time.Millisecond)
}
