/**
 * Goのプログラムは、エラーの状態を error 型で表現する
 * error 型は組み込みのインタフェースで、以下のように定義されている
 *   type error interface {
 *       Error() string
 *   }
 *
 * エラーが起こりうる関数は基本的に error 変数を返す
 * 呼び出し元はエラーが nil かどうかを確認することでエラーをハンドル(取り扱い)する
 *   i, err := strconv.Atoi("42") // 文字列を整数に変換 => 変換後の値と error 値を返す
 *   if err != nil {
 *       // エラーが起こった場合の処理
 *       fmt.Printf("couldn't convert number: %v\n", err)
 *       return
 *   }
 *   fmt.Println("Converted integer:", i)
 */

package main

import (
	"fmt"
	"time"
)

// 新しいエラー型: MyError
type MyError struct {
	When time.Time // いつエラーが起こったか
	What string // エラー文
}
// エラー型は Error() string メソッドを持つ必要がある
func (e *MyError) Error() string {
	// いつ、どんなエラーが起きたかを文字列として返す
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// run 関数: わざとエラーを返す
func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func main() {
	if err := run(); err != nil {
		// エラーが起きた場合、そのエラー値を出力
		fmt.Println(err)
	}
}
