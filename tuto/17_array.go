/**
 * T 型の変数を n 個もつ配列は [n]T で表現される
 * 配列 a の i 番目の変数にアクセスする際は a[i] でアクセスする
 */

package main

import "fmt"

func main() {
	// string型 2つからなる配列 a を宣言
	var a [2]string
	// Goの配列のインデックスは 0 番からはじまるため、配列 a のインデックスは 0, 1 の2つがある
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
	
	// 配列を初期化子つきで宣言する場合は {<初期値>} ステートメントを使う
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 配列の長さは len 関数で取得できる
	fmt.Println(len(primes))
}
