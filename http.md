```go
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

var dir string
var port int
var staticHandler http.Handler

// 初始化参数
func init() {
	dir = path.Dir(os.Args[0])
	flag.IntVar(&port, "port", 8990, "服务器端口")
	flag.Parse()
	staticHandler = http.FileServer(http.Dir(dir))
}
func main() {
	http.HandleFunc("/", StaticServer)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		staticHandler.ServeHTTP(w, req)
		return
	}
	io.WriteString(w, "hello, world!\n")
}
```