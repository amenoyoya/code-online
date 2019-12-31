/**
 * Goでは最初の文字が大文字で始まる名前が、外部パッケージから参照可能な名前である
 * 例えば Pi は math パッケージでエクスポートされているが pi はエクスポートされていない
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// math.pi は undefined エラーになる
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
