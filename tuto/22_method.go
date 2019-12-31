/**
 * Goに class の概念は存在しない
 * 代わりに、構造体に対してメソッドを定義することができる
 *   func (<インスタンス名> <構造体型>) <メソッド名> ([引数]) <戻り値型> { ... }
 * ただし、同じパッケージ内で定義される必要がある
 * func と メソッド名 の間に記述する引数（<インスタンス名> <構造体型>）をレシーバ引数と呼ぶ
 */

package main

import (
	"fmt"
	"math"
)

// フィールド X, Y を持つ Vertex 構造体
type Vertex struct {
	X, Y float64
}

// Vertex 構造体に Abs メソッドを定義
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 構造体のフィールド変数を変更するメソッドを定義する場合は、レシーバ引数にポインタを指定する必要がある
// （引数に渡される構造体は参照ではなくコピーであるため）
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/**
 * メソッドは構造体以外の任意の型（type で定義される型）に対しても定義できる
 */
// float64型の別名で Float 型を定義
type Float float64

// Float 型に Abs メソッドを定義
func (f Float) Abs() Float {
	if f < 0 {
		return -f
	}
	return f
}

func main() {
	// X: 3, Y: 4 の Vertex インスタンス化
	v := Vertex{3, 4}
	// メソッド呼び出し: <構造体インスタンス>.<メソッド名>([引数])
	fmt.Println(v.Abs())
	// v の X, Y を 2倍する
	v.Scale(2)
	fmt.Println(v)

	// -0.1 の Float 型
	f := Float(-0.1)
	fmt.Println(f.Abs())
}
