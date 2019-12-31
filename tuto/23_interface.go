/**
 * 同じメソッド名を持つ別の構造体を同じように扱うための機能として Interface がある
 *   type <インタフェース名> interface {
 *       <メソッド名>([引数]) <戻り値型>
 *       ...
 *   }
 *
 * メソッドを一つも持たないインタフェースは空のインタフェースと呼ばれ、任意の値を保持することができる
 *   type Any interface {}
 */

package main

import "fmt"

// Greeter インタフェース: Hello メソッドを持つ
type Greeter interface {
	Hello()
}

// hello関数: 引数に渡された構造体の Hello メソッドを実行
func hello(greeter Greeter) {
	greeter.Hello()
}

// Human 構造体: Hello メソッドは "私は {name} です" と出力
type Human struct {
	Name string
}
func (human Human) Hello() {
	fmt.Printf("私は %s です\n", human.Name)
}

// Cat 構造体: Hello メソッドは "ニャー" と出力
type Cat struct {}
func (cat Cat) Hello() {
	fmt.Println("ニャー")
}

func main() {
	// Greeter インタフェースは、Hello メソッドを持つ任意の構造体をインスタンス化できる
	var greeter Greeter = Human{"yoya"} // Name: "yoya" の Human 構造体を Greeter インタフェースを通じてインスタンス化
	// greeter.Hello() を実行
	hello(greeter)

	// greeter に Cat 構造体のインスタンスを代入
	greeter = Cat{}
	// greeter.Hello() を実行
	hello(greeter)
}
