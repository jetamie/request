### Golang封装请求(带cookie)

>安装方法

```shell
go get github.com/jetamie/request
```

>使用方法

```go
package request

import (
	"testing"
)

func TestRequest_SetHost(t *testing.T) {
	r := NewRequest()
	//配置域名host访问
	r.SetHost("www.baidu.com").Request("http://112.80.248.75:80")

}

func TestRequest_SetLog(t *testing.T) {
	r := NewRequest()
	//打印相关请求信息，默认false
	r.SetLog(true).Request("https://www.baidu.com")
}

func TestRequest_SetPostData(t *testing.T) {
	r := NewRequest()
	//发送post数据
	post := map[string]string{}
	post["username"] = "xxx"
	post["password"] = "xxx"
	r.SetPostData(post).Request("https://passport.baidu.com/v2/api/?login")
}

func TestRequest_Request(t *testing.T) {
	//普通GET请求
	r := NewRequest()
	r.Request("https://www.baidu.com")
	//t.Logf("body:%s\r\n", body)
}

func TestRequest_PrintCookie(t *testing.T) {
	//打印cookies信息
	r := NewRequest()
	r.Request("https://www.baidu.com")
	r.PrintCookie()
}
```

>版本介绍

`v0.1`

- get/post请求(带cookie)
- 请求https时跳过ssl证书验证
- 允许动态设置Host访问
- 允许获取当前用户请求的cookie

---

_交流:jetamiett@163.com_
