## 小程序二维码服务

使用 golang 

> 简单、高效、跨平台、无需运行环境


## 分发机制

我们为每一个平台 提供可执行程序


## 交叉编译


```
gox -os=\"linux\"
```


## 短网址 [API 说明](http://dwz.cn/api)

在 v1.0.3 版本开始 使用 百度的 短网址服务，

解决小程序 path 路径限制长度问题


## 显示原网址

### 请求

向 `http://dwz.cn/query.php` 发送 `post` 请求，

发送数据包

```
tinyurl=查询的短地址
```

### 返回

json格式的数据

```json
{
 "status":"0", // !=0 出错
 "err_msg":"", // 获得错误信息（UTF-8编码）
 "longurl":"" // 返回原网址
}
```

## 生成短网址

## 请求

向 `http://dwz.cn/create.php` 发送 `post` 请求

发送数据包
```
url=长网址
```

## 返回

json格式的数据

```json
{
 "status":"0", // !=0 出错
 "err_msg":"", // 获得错误信息（UTF-8编码）
 "tinyurl":"" // 生成的短网址
}
```





## 静态编译

```
$ go get github.com/jteeuwen/go-bindata/...
$ go get github.com/elazarl/go-bindata-assetfs/...
$ go-bindata-assetfs html/...
```

main.go

```go
http.Handle("/", http.FileServer(assetFS()))
```

## 在线demo

https://appcode2.anla.io/
