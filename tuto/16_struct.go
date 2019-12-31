/**
 * struct（構造体）は複数の変数をまとめた、フィールド集合体であり、以下のように宣言される
 *   type <構造体名> struct { <フィールド> }
 *
 * フィールドに値を入れて構造体を実体化させる場合は <構造体名>{<フィールド値>} でインスタンス化する
 */

package main

import "fmt"

// 変数 X, Y をもつ Vertex 構造体
type Vertex struct {
	X int
	Y int
}

func main() {
	// X: 1, Y: 2 の Vertex構造体をインスタンス化
	v := Vertex{1, 2}
	fmt.Println(v)

	// v のフィールド X に 4 を代入
	v.X = 4
	fmt.Println(v)

	// 構造体のポインタ
	p := &v

	// ポインタの示す先にある構造体のフィールドにアクセスする場合 (*p).X のように書く
	// ただし Go では、上記表記法は面倒であるため p.X と書いても同一の挙動をするように設計されている
	(*p).X = 1e3 // ポインタ p を通して v.X に 1000 を代入
	p.Y = 1e2 // ポインタ p を通して v.Y に 100 を代入
	fmt.Println(v)

	// var は () をつけると複数行で宣言できる
	var (
		// 構造体インスタンス化の際、特定のフィールドのみに値を割り当てることも可能
		v1 = Vertex{X: 1} // X: 1, Y: ゼロ値 の Vertex をインスタンス化
		v2 = Vertex{} // X: ゼロ値, Y: ゼロ値 の Vertex をインスタンス化
		pv = &Vertex{1, 2} // X: 1, Y: 2 の Vertex インスタンスへのポインタ
	)
	fmt.Println(v1, v2, pv)
}
