/**
 * select ステートメントは、複数の case のいずれかが準備完了するまでチャネルをブロックする
 * 複数の case が準備完了になった場合は、ランダムに選択される
 *   select {
 *   case [<変数> :=] <- <チャネル> :
 *       <指定チャネルが準備完了した時の処理>
 *   [default: <どのチャネルも準備完了していない時の処理>]
 *   }
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Tick(n): n秒ごとの実行プロセスを作成
	tick := time.Tick(100 * time.Millisecond) // 0.1秒ごとの処理

	// time.After(n): n秒後の実行プロセスを作成
	boom := time.After(500 * time.Millisecond) // 0.5秒後の処理

	// 無限ループ
	for {
		select {
		case <-tick:
			// 0.1秒ごとに "tick." を出力
			fmt.Println("tick.")
		case <-boom:
			// 0.5秒後に "BOOM!" を出力してループ終了
			fmt.Println("BOOM!")
			return
		default:
			// 上記のいずれのチャネルも準備されていない場合は "    ." を出力して 0.05秒待機
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
