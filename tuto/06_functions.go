/**
 * Goの関数定義は
 *   func <関数名>([<引数> <引数型名>]) <戻り値型名> {
 *        <関数内部処理>
 *   }
 * という形で宣言される
 */

package main

import "fmt"

// add関数: 2引数 x, y をとり x + y の計算結果を返す
// ※ 関数の2つ以上の引数が同じ型である場合は最後の型を残して省略できる
//    func add(x, y int) int {...}
func add(x int, y int) int {
	return x + y
}

// 関数は複数の戻り値を返すことも可能
// swap関数: 2引数 x, y をとり y, x を返す
func swap(x, y string) (string, string) {
	return y, x
}

// Goでは戻り値となる変数に名前をつけることも可能
// split関数: 1引数 sum をとり x(= sum * 4 / 9), y(= sum - x) を返す
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 戻り値のある関数の最後には必ず return 文を入れる必要がある
}

func main() {
	fmt.Println(add(42, 13)) // => 42 + 13 = 55 が表示される
	fmt.Println(swap("Hello", "World")) // "World", "Hello" の順で表示される
	fmt.Println(split(17)) // 7, 10 が表示される
}
