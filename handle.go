package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// CodeURL is ...
func CodeURL(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	path := req.FormValue("path")
	width := req.FormValue("width")

	if path == "" {
		io.WriteString(w, "path not find")
		return
	}

	if width == "" {
		width = "200"
	}

	var qcode QuCode

	shortUl := CreatBaiduShortURL(path)

	if shortUl == "" {
		io.WriteString(w, "地址错误")
		return
	}

	qcode.Path = Config.WeiXin.Path + shortUl

	qcode.Width = width

	token := GetToken()

	body := Createwxaqrcode(token, qcode)

	io.WriteString(w, body)
}

// PathRandom is ...
func PathRandom(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	path := req.FormValue("path")
	width := req.FormValue("width")

	if path == "" {
		io.WriteString(w, "path not find")
		return
	}

	if width == "" {
		width = "200"
	}

	var qcode QuCode

	qcode.Path = path

	qcode.Width = width

	token := GetToken()

	body := Createwxaqrcode(token, qcode)

	io.WriteString(w, body)

}

// APIInfo is ...
func APIInfo(w http.ResponseWriter, req *http.Request) {

	data, _ := json.Marshal(Config.APPInfo)
	w.Write([]byte(data))
}

var staticHandler http.Handler

// StaticServer is 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/code" {
		staticHandler.ServeHTTP(w, req)
		return
	}
}
