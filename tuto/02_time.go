/*
 * 現在日時を取得するプログラム
 */

package main

// 複数パッケージをインポートする場合、以下のような書き方ができる
import (
	"fmt"
	"time"
)

// エントリーポイント
func main() {
	// "Welcome to the playground!" を表示
	fmt.Println("Welcome to the playground!")

	// time.Now(): 現在日時を取得
	fmt.Println("The time is", time.Now())
}
