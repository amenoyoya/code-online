/**
 * Goでは var ステートメントで変数を宣言できる
 *   var <変数名> <型名> [= <初期値>]
 *
 * Go言語の基本型（組み込み型）は以下の通り
 * - bool (true/false)
 * - string (文字列型)
 * - int, int8, int16, int32, int64 (整数型)
 * - uint, uint8, uint16, uint32, uint64 (符号なし整数型)
 * - byte (uint8 の別名 = 8ビット符号なし整数型)
 * - rune (int32 の別名 = 32ビット整数型: Unicodeのコードポイントを表す)
 * - float32, float64 (浮動小数点型)
 * - complex64, complex128 (複素数型)
 */

package main

import "fmt"

// bool型変数 c, python, java の宣言
// 初期値を与えずに宣言するとゼロ値が与えられる（0, false, "" など）
var c, python, java bool

func main() {
	// int型変数 i: main関数の中でのみ使用可能
	var i int
	// c, python, java は main関数の外で宣言されているが、main関数内部で参照可能
	fmt.Println(i, c, python, java)

	// 変数宣言時に初期化子（=）により値を代入できる
	var x, y int = 1, 2
	// 初期化子が与えられている場合、型は省略しても良い
	var z = "ゼット"

	fmt.Println(x, y, z)
}
