/**
 * var ステートメントによる変数宣言の短縮形で := による初期化子が使える（暗黙的に型宣言される）
 * ただしこれは関数内部でのみ使用可能であり、関数の外では var, func 等の宣言キーワードが必須になる
 */

package main

import "fmt"

func main() {
	// 以下は i, j int = 1, 2 と同等
	i, j := 1, 2

	fmt.Println(i, j)
}
