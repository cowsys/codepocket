package main

import "fmt"
import "time"

// ref:http://go-talks.appspot.com/github.com/lestrrat/go-slides/2014-yapcasia-go-for-perl-mongers/main.slide#58
func main() {
	fmt.Println("start")
	defer fmt.Println("end")
	// スタンダードな同期パターン
	// 1.呼び出し元で終了判定のためのchannelを定義し
	// 2.呼び出すgoroutineに渡し、その中で終了したら終了判定を設定
	// 3.呼び出し元の"<-done"でgoroutineの実行終了を待つ
	// 以上の3ステップで整合性を持って同期できる

	// NOTE:メイン関数先に終了する問題に対処できる
	done := make(chan struct{})

	go func() {
		defer func() {
			done <- struct{}{}
		}()

		time.Sleep(5 * time.Second)
		fmt.Println("Goroutine fired!!")
	}()

	<-done
}
