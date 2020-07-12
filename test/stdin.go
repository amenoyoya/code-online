/*
標準入力から文字列取得

# 一行入力
$ echo 'Hello, world' | go run test/stdin.go

# 複数行入力
$ go run test/stdin.go << EOS
1. test
2. stdin
EOS

# JSON入力
$ go run test/stdin.go << EOS
{
    "name": "Alice",
    "sex": "female",
    "age": 13
}
EOS
*/

package main

import (
    "bufio"
    "os"
    "log"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

func main() {
    // 標準入力 Reader
    reader := bufio.NewReader(os.Stdin)
    // 先頭が "{" | "[" なら JSON として parse
    head, _ := reader.Peek(1)
    if string(head) == "{" || string(head) == "[" {
        var data interface{}
        dec := json.NewDecoder(reader)
        if err := dec.Decode(&data); err != nil {
            // エラー時: 標準エラーに出力してプログラム終了
            log.Fatal(err)
        }
        fmt.Println("json")
        fmt.Println(data)
    } else {
        // ioutil.ReadAll で標準入力を全て読み込み
        buf, err := ioutil.ReadAll(reader)
        if err != nil {
            // エラー時: 標準エラーに出力してプログラム終了
            log.Fatal(err)
        }
        fmt.Println("text")
        fmt.Println(string(buf))
    }
}
