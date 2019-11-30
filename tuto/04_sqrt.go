/**
 * 平方根の計算と書式付き文字列表示
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	 * fmt.Printf 関数は以下のような書式を使用することができる
	 * 	論理値(bool): %t
	 * 	符号付き整数(int, int8など): %d
	 * 	符号なし整数(uint, uint8など): %d
	 * 	浮動小数点数(float64など): %g
	 * 	複素数(complex128など): %g
	 * 	文字列(string): %s
	 * 	チャネル(chan): %p
	 * 	ポインタ(pointer): %p
	 * 	デフォルト書式: %v
	 * 第2引数以降に指定された値が、上記で指定された書式に変換されて表示される
	 * なお、fmt.Printf は自動的に改行しないため、文字列の末尾に改行文字 '\n' を指定したほうが良い
	 */
	// math.Sqrt(x float64) float64: xの平方根を返す
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
