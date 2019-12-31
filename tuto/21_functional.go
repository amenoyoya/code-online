/**
 * Goでは関数も変数として扱われ、無名関数（関数名を持たない関数: func(<引数>){...}）も使える
 * 関数値は、関数の引数に取ることも、戻り値として利用することも可能
 */

package main

import (
	"fmt"
	"math"
)

// compute関数: float64型引数を2つとりfloat64型の値を返す関数を引数にとる
//   引数に指定された関数に対して、引数 3, 4 を与えて実行する
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder関数: ローカル変数 sum に x を加算する関数を生成
func adder() func(int) int {
	sum := 0
	// Goの関数はクロージャであるため関数の外部で宣言された変数をキャプチャなしで参照できる
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// 変数 hypot に無名関数を代入
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// √5^2 + 12^2 = 13
	fmt.Println(hypot(5, 12))
	// √3^2 + 4^2 = 5
	fmt.Println(compute(hypot))
	// 3^4 = 81
	fmt.Println(compute(math.Pow))

	// pos 関数と neg 関数を生成
	pos, neg := adder(), adder()
	// 10回ループ
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // 0 + 1 + 2 + ... + 9
			neg(-2 * i), // 0 + (-2) + (-4) + ... + (-18)
		)
	}
}
