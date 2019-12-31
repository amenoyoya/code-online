/**
 * 他の言語における連想配列（辞書配列）と同じものがGoにも存在し、map型として定義される
 *   map[<キーの型>]<値の型>
 *
 * マップはキーと値を関連づけるものであり、そのゼロ値は nil である
 * マップを作成する場合は make 関数を使うか、初期化子 {<キー>: <値>, ...} を用いて宣言する
 */

package main

import "fmt"

// float64型の Lat, Long フィールドをもつ Vertex 構造体
type Vertex struct {
	Lat, Long float64
}

func main() {
	// キーが文字列で値がVertex型であるマップを作成
	m := make(map[string]Vertex)

	// m のキー "Bell Labs" に Vertex{40.68433, -74.39967} を代入
	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	fmt.Println(m)

	// 初期化子つきでマップを作成
	m2 := map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google": Vertex{37.42202, -122.08408},
	}
	fmt.Println(m2)

	// マップのキーと関連する値を削除する場合は delete 関数を使う
	//   delete(<マップ変数>, <削除対象キー>)
	delete(m2, "Bell Labs")
	fmt.Println(m2)

	// 特定のキーが存在するかどうかは、要素取得時の第2戻り値で確認可能
	_, ok := m2["Bell Labs"]
	fmt.Println("m2 has key 'Bell Labs':", ok)
}
