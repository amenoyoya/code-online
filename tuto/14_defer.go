/**
 * defer ステートメントにより関数を遅延実行することができる
 * defer へ渡した関数の引数はすぐに評価されるが、その関数自体は呼び出し元関数が return するまで実行されない
 */

package main

import "fmt"

func hello() {
	// 以下の関数は hello関数終了時に呼びされる
	// => "hello" "world" の順に表示される
	defer fmt.Println("world")
	
	// 以下の関数が先に実行される
	fmt.Println("hello")
}

func count() {
	// defer へ渡した関数が複数ある場合、それらはスタックへ積まれ、LIFO(last-in-first-out)の順で実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("counting")
}

func main() {
	hello()
	count()
}
