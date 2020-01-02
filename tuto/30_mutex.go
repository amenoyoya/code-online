/**
 * goroutine でデータ同期をするためにチャネルは有用だが、データの排他制御を行いたい場合は sync.Mutex ライブラリを使うと良い
 *   func (sync.Mutex) Lock(): データをロックし、排他制御を開始する
 *   func (sync.Mutex) Unlock(): データをアンロックし、データへの書き込みを可能にする
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter: 任意のキー文字列に紐づくカウントを排他的に制御可能な構造体
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex // 排他制御用ミューテックス
}

// SafeCounter.Inc メソッド: 指定文字列に紐づくカウントを進める
// SafeCounter のフィールドに変更を加えるため、レシーバ引数は SafeCounter のポインタにする
func (counter *SafeCounter) Inc(key string) {
	counter.mux.Lock()   // 並列実行される他のプロセスから書き込みされるのを防ぐため、データロック
	counter.v[key]++     // 指定キーのカウントをインクリメント
	counter.mux.Unlock() // 排他制御を終了
}

// SafeCounter.Value メソッド: 指定文字列に紐づくカウント値を取得
func (counter *SafeCounter) Value(key string) int {
	counter.mux.Lock()
	// カウント値を返した後にデータをアンロックするため defer で遅延実行
	defer counter.mux.Unlock()
	// カウント値を返す
	return counter.v[key]
}

func main() {
	counter := SafeCounter{v: make(map[string]int)}
	// "somekey" に紐づくカウントを1000回分進める（並列実行）
	for i := 0; i < 1000; i++ {
		go counter.Inc("somekey")
	}
	// 1秒待機
	time.Sleep(time.Second)
	// 1秒待機した後、並列実行されていた Inc メソッドにより、どこまでカウント値が進んでいるか確認
	fmt.Println(counter.Value("somekey"))
}
