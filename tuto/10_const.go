/**
 * Goでは const キーワードを用いて定数を宣言できる
 * 定数は変数と違い、別の値を再代入することはできない
 * また定数として宣言できるのは string, bool, 数値(int, float64, ...etc) のみである
 */

package main

import "fmt"

// 定数 Pi: 3.14
const Pi = 3.14

func main() {
	// 定数に別の値を代入することはできない（cannot assign to Pi）
	// Pi = 4.13
	fmt.Println("Happy", Pi, "Day")
}
