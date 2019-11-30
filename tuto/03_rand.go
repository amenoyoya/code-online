/**
 * 乱数表示プログラム
 */
package main

import (
	"fmt"
	"time"
	"math/rand" // mathモジュール内のrandモジュール利用
)

func main() {
	// 乱数の種はデフォルトでは 1 で固定となっているため、何回プログラムを実行しても同じ値が返ってくる
	// => 変更する場合は rand.Seed(seed int64) で任意の種を与える必要がある
	rand.Seed(time.Now().UnixNano()) // 現在日時をナノ秒に変換した値を種にする

	// rand.Intn(n int) int: 0..n-1 の整数乱数を返す
	fmt.Println("My favorite number is", rand.Intn(10))
}
