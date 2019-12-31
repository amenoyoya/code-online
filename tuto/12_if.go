/**
 * 条件分岐は if, else if, else 構文で記述する
 *   if [評価ステートメント;] <条件式> {
 *       <条件合致時の処理>
 *   }
 *   [else if <条件式> { ... }]
 *   [else { ... }]
 *
 * 評価ステートメントでは、条件式用の変数宣言を行うことができる（if, else if, else 式の内部でのみ使用可能な変数）
 */

package main

import (
	"fmt"
	"math"
)

// pow関数: x^n を計算し、それが lim1 未満なら x^n を返し、lim1 以上 lim2 未満なら lim1 を返し、lim2 以上なら lim2 を返す
func pow(x, n, lim1, lim2 float64) float64 {
	if v := math.Pow(x, n); v < lim1 {
		return v
	} else if lim1 <= v && v < lim2 {
		return lim1
	} else {
		return lim2
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 18, 40),
		pow(3, 3, 18, 40),
		pow(3, 4, 18, 40),
	)
}
