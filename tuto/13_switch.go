/**
 * switch - case ステートメントを用いて if - else 文を短く書くことができる
 *   switch [評価ステートメント;] <評価対象変数> {
 *   case <合致対象> :
 *       <条件合致時の処理>
 *   [default: <条件に合致しない場合の処理>]
 *   }
 *
 * Goの switch 文では、他の言語の switch と異なり、該当した case のみ実行され、case も定数である必要はない
 * また、評価対象変数も数値である必要はない
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	// 実行OSを判定
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	// switch の評価対象を省略すると switch true と同一のステートメントになる
	// この構文は冗長になりがちな if - else if - else 文をシンプルに表現するために用いられる
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("おはよう")
	case t.Hour() < 17:
		fmt.Println("こんにちは")
	default:
		fmt.Println("こんばんは")
	}
}
