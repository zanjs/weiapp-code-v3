package main

// Config is ...
var Config = struct {
	APPInfo struct {
		Name    string `default:"海淘助手小六哦"`
		Version string `default:"1.0.9"`
		Port    int    `default:"12300"`
		Contact string `default:"anlasheng@gmail.com By Julian -anla"`
	}
	WeiXin struct {
		Appid     string `default:""`
		Secret    string `default:""`
		GrantType string `default:"client_credential"`
		Path      string `default:"page/scanCode/detail/detail?code="`
		TokenURL  string `default:"https://api.weixin.qq.com/cgi-bin/token"`
		CreateURL string `default:"https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="`
	}
	QrCode struct {
		DefaultContent string `default:"https://m.6city.com"`
	}
	BaiDu struct {
		CreateURL string `default:"http://dwz.cn/create.php"`
	}
	About []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

// QuCode is ...
type QuCode struct {
	Path  string `json:"path"`
	Width string `json:"width"`
}

// ShortU is ...
type ShortU struct {
	URL string `json:"url"`
}
