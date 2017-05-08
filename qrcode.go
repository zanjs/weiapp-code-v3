package main

import (
	"io"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
)

// CreateUQrcode is
func CreateUQrcode(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	sURL := req.FormValue("url")

	if sURL == "" {
		io.WriteString(w, "url not find")
		return
	}
	var png []byte
	png, err := qrcode.Encode(sURL, qrcode.Medium, 300)

	if err != nil {

	}

	io.WriteString(w, string(png))

	// _ := qrcode.WriteFile(sURL, qrcode.Medium, 256, "qr.png")

}

// Qcode is ..
func Qcode(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	_content := req.FormValue("content")

	if _content == "" {
		_content = Config.QrCode.DefaultContent
	}

	var png []byte

	png, _ = qrcode.Encode(_content, qrcode.Medium, 256)

	io.WriteString(w, string(png))
}
