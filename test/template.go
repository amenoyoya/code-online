/**
 * go text/template によるテンプレート処理
 *
 * # 標準出力に処理結果表示
 * $ go run test/template.go
 *
 * # 文字列に処理結果を入れてから標準出力に表示
 * $ go run test/template.go -mode string
 *
 * # ファイルに処理結果を書き込み
 * $ go run test/template.go -mode <ファイル名>
 */
package main

import (
    "log"
    "os"
    "fmt"
    "bytes"
    "flag"
    "text/template"
)

/** テンプレート
 * {{- ... -}}: 直前と直後の改行と空白を全て削除
 * {{- ... }}: 直前の改行と空白を全て削除
 * {{ ... -}}: 直後の改行と空白を全て削除
 */
var tmpl string = `
server {
    listen 80;
    server_name {{.Host}};
{{- if .CertFile }}
    return 301 https://{{.Host}}$request_uri;
{{- else}}
    location / {
{{- if .Root}}
        root {{.Root}};
        index index.html;
{{- else if .Port}}
        proxy_set_header Host $http_host;
        proxy_pass http://localhost:{{.Port}};
{{- end}}
{{- if .BasicAuth}}
        auth_basic "StaffOnly";
        auth_basic_user_file {{.BasicAuth}};
{{- end}}
    }
{{- end}}
}

{{- if .CertFile}}
server {
    listen 443 ssl;
    server_name {{.Host}};
    location / {
{{- if .Root}}
        root {{.Root}};
        index index.html;
{{- else if .Port}}
        proxy_set_header Host $http_host;
        proxy_pass http://localhost:{{.Port}};
{{- end}}
{{- if .BasicAuth}}
        auth_basic "StaffOnly";
        auth_basic_user_file {{.BasicAuth}};
{{- end}}
    }
    ssl_certificate {{.CertFile}};
    ssl_certificate_key {{.CertKey}};
}
{{- end}}
`

// NginxConf構造体: 構造体のフィールド名は大文字で始める必要がある
type NginxConf struct {
    Host string // server_name
    Port int // port_forwarding
    Root string // root_dir
    CertFile string // ssl_certificate
    CertKey string // ssl_certificate_key
    BasicAuth string // auth_basic_user_file
}

func main() {
    conf := NginxConf{
        "example.jp", 0, "/var/www/html/", "/etc/lego/example.jp.crt", "/etc/lego/example.jp.key", "",
    }
    // template.New(<テンプレート名>).Parse(<文字列>)
    t, err := template.New("nginx.conf").Parse(tmpl)
    if err != nil {
        log.Fatal(err)
    }
    // オプション引数処理
    var (
        mode = flag.String("mode", "stdout", "処理モード")
    )
    flag.Parse()
    fmt.Println(*mode)
    // 標準出力
    if *mode == "stdout" {
        // Execute(io.Writer(出力先), <データ>)
        if err = t.Execute(os.Stdout, conf); err != nil {
            log.Fatal(err)
        }
    } else if *mode == "string" {
        // 文字列に出力
        var res bytes.Buffer
        if err = t.Execute(&res, conf); err != nil {
            log.Fatal(err)
        }
        fmt.Println(res.String())
    } else {
        // ファイルに出力
        fp, err := os.OpenFile(*mode, os.O_RDWR|os.O_CREATE, 0644)
        defer fp.Close() // ブロック終了時に fp.Close() 呼び出し
        if err != nil {
            log.Fatal(err)
        }
        if err = t.Execute(fp, conf); err != nil {
            log.Fatal(err)
        }
        fmt.Println(*mode, "に書き込みました")
    }
}
