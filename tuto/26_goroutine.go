/**
 * Goでは並列処理を行うための機構として goroutine という軽量なスレッドシステムを持っている
 *   go f(x, y, z)
 * と記述すれば新しいスレッドが起動し、メイン処理と並列的に f(x, y, z) 関数が実行される
 * ただし、goroutine は同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期を取る必要がある
 */

package main

import (
	"fmt"
	"time"
)

// say 関数: 引数に与えられた文字列を0.1秒毎に5回出力する
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// say 関数を並列実行
	go say("world")
	
	// 上記とは別に say 関数を実行
	say("hello")
}
