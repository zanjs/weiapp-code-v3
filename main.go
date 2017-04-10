package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/configor"
	// qrcode "github.com/skip2/go-qrcode"
)

var port int

// 初始化参数
func init() {

	configor.Load(&Config, "config.yml")

	var welcomeStr = "\n欢迎使用 微信小程序二维码生成 微服务 \n- 高性能、跨平台、易扩展 \n- Version: " + Config.APPInfo.Version
	var portRed = "\n- 服务占用 " + strconv.Itoa(Config.APPInfo.Port) + "端口号"
	var ymlStr = "\n\nconfig.yml 文件为自定义配置文件, 您可以根据场景修改它"
	var openStr = welcomeStr + portRed + ymlStr + "\n\n请打开浏览器输入 \nhttp://localhost:"

	flag.IntVar(&port, "port", Config.APPInfo.Port, "服务器端口")
	fmt.Println(openStr + fmt.Sprintf("%d", port))
	flag.Parse()

}

func main() {

	// http.HandleFunc("/html/", StaticServer)
	http.Handle("/", http.FileServer(http.Dir("./html")))
	// http.Handle("/", http.FileServer(assetFS()))

	http.HandleFunc("/code", CodeURL)
	http.HandleFunc("/rand", PathRandom)
	http.HandleFunc("/api", APIInfo)
	http.HandleFunc("/qrcode", CreateUQrcode)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

var (
	logFileName = flag.String("log", "log.log", "Log file name")
)
