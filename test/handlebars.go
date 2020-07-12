/*
Handlebars Template Engine によるテキスト展開
$ go run test/handlebars.go
*/
package main

import (
    "fmt"
    "github.com/aymerick/raymond"
)

var src string = `
{{title}}
{{#if body}}
---
{{body}}
{{/if}}
`

func main() {
    res, err := raymond.Render(src, map[string]interface{} {
        "title": "My Handlebars Title",
        "body": "Hello Handlebars by raymond",
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
