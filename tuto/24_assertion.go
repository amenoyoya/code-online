/**
 * 型アサーションは、インターフェースの値の基になる具体的な値を利用する手段を提供する
 * 例えば、あるインタフェースの値 i が、具体的な型 T の値を保持している場合に、その値を変数 t に代入したい場合
 *   t := i.(T)
 * と記述する
 * このとき、i が T 型の値を保持していない場合、この文は panic を起こす
 *
 * なお、型アサーションは第2戻り値に、型を保持しているかどうかの真偽値を格納する
 *   t, ok := i.(T)
 */

package main

import "fmt"

func main() {
	// 任意型（空のインタフェース）の変数 i に "hello" という文字列を代入
	var i interface{} = "hello"
	
	// i が文字列型の値を保持している場合、変数 s にその値を代入
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// i が浮動小数点型の値を保持している場合、変数 f にその値を代入
	// 変数 ok には、浮動小数点型を保持しているかどうかの真偽値を代入
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 以下の文は panic を引き起こす
	// panic が起こると以降の処理は実行されない
	//f = i.(float64)
	//fmt.Println(f)

	// 型アサーションは switch 文と組み合わせることも可能
	switch value := i.(type) {
	case string:
		fmt.Printf("\"%s\" は文字列型です\n", value)
	case int:
		fmt.Printf("%d は整数型です\n", value)
	case float64:
		fmt.Printf("%f は浮動小数点型です\n", value)
	default:
		fmt.Println(value, "は知らない型です")
	}
}
