/**
 * for ループでは range キーワードでスライスやマップ等の要素を逐次反復処理することができる
 *   for <インデックス変数>, <要素変数> := range <スライス等> { <反復処理> }
 */

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64}

func main() {
	// 以下の for ループは
	//   for i := 0; i < len(pow); i++ { v := pow[i]; .... }
	// と等価
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	// 使わない値は _ へ代入することで捨てることも可能
	// 以下はインデックスが不要で、スライスの各要素の値のみ必要な場合の処理
	for _, v := range pow {
		fmt.Println(v)
	}
}
