/**
 * Goの for ループは以下のような構文となる
 *   for [初期化ステートメント]; [条件式]; [後処理ステートメント] { <ループ処理> }
 *
 * まず初期化ステートメントが実行され、条件式が true なら最初のループ処理が実行される
 * その後、後処理ステートメントが実行され、条件式が true なら2回目のループ処理が実行される
 * 条件式が false になるまで、後処理ステートメント => ループ処理 が繰り返し実行される
 */

package main

import "fmt"

func main() {
	sum := 0
	// i = 0 から始まり i < 10 の条件に合致する限り sum に1を加算し続ける（i は1ずつ加算される）
	// => 10回ループ
	for i := 0; i < 10; i++ {
		sum += 1 // sum++ でも可
	}
	fmt.Println(sum)

	// Goに while 文はなく、初期化ステートメントと後処理ステートメントを省略した for 文で代替される
	// ※条件式も省略すると無限ループになる
	sum = 1
	for sum < 1000 {
		sum += sum // sum に sum を加算 => sum が倍々になる
	}
	fmt.Println(sum)
}
