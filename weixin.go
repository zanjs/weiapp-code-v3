package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GetToken is ...
func GetToken() string {
	u, _ := url.Parse(Config.WeiXin.TokenURL)
	q := u.Query()
	q.Set("grant_type", Config.WeiXin.GrantType)
	q.Set("appid", Config.WeiXin.Appid)
	q.Set("secret", Config.WeiXin.Secret)
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	resp, err := http.Get(u.String())

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	byt := []byte(body)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	accessToken := dat["access_token"].(string)

	return accessToken
}

// Createwxaqrcode is...
func Createwxaqrcode(token string, qcode QuCode) string {

	creatURL := Config.WeiXin.CreateURL

	strJSON := `{"path": "page/my/my", "width": 430}`

	//struct 到json str
	if b, err := json.Marshal(qcode); err == nil {

		strJSON = string(b)
	}

	creatURL += token
	resp, err := http.Post(creatURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(strJSON))

	if err != nil {

	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
