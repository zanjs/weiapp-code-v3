package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// CreatBaiduShortURL is ...
func CreatBaiduShortURL(logURL string) string {
	creatURL := Config.BaiDu.CreateURL

	strJSON := `{"url": ""}`

	var shortu ShortU

	shortu.URL = logURL
	//struct 到json str
	if b, err := json.Marshal(shortu); err == nil {

		strJSON = string(b)
	}

	// fmt.Println(strJSON) //创建 短网址 Body

	resp, err := http.Post(creatURL,
		"application/x-www-form-urlencoded",
		strings.NewReader("url="+logURL))

	if err != nil {

	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	byt := []byte(body)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	// fmt.Println(dat) // 短网址 返回信息

	// longurl := dat["longurl"].(string)

	// 定义一个文件
	// fileName, err := os.Open("log.log")

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)

	//write log
	log.Printf(":%v : logurl: %v \n", dat, strJSON)

	// debugLog.Println("A debug message here")
	//配置一个日志格式的前缀
	// debugLog.SetPrefix("[Info short :]")
	// debugLog.Println(dat)
	//配置log的Flag参数
	// debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	// debugLog.Println("A different prefix")

	status := dat["status"].(float64)

	// fmt.Println(reflect.TypeOf(status)) // type

	if status == 0 {
		accessURL := dat["tinyurl"].(string)
		// fmt.Println(accessURL)
		return string(accessURL)
	}

	nourl := ""

	return string(nourl)

}
