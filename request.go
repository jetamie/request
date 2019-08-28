package request

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

type Request struct {
	CurCookies []*http.Cookie
	CurCookieJar *cookiejar.Jar
	host string
	log bool
	post map[string]string
}

func NewRequest() *Request {
	curCookieJar,_ := cookiejar.New(nil)
	return &Request{
		CurCookies:nil,
		CurCookieJar:curCookieJar,
		host:"",
		log:false,
		post:nil,
	}
}

func (r *Request) SetHost(host string) *Request {
	r.host = host
	return r
}

func (r *Request) SetLog(logged bool) *Request {
	r.log = logged
	return r
}

func (r *Request) SetPostData(post map[string]string) *Request {
	r.post = post
	return r
}

func (r *Request) Request(urls string) string {
	if r.log {//输出日志
		log.Println("[GET] " + urls)
	}
	client := &http.Client{
		Jar:r.CurCookieJar,
		Transport:&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	//拼接请求
	var req *http.Request
	if r.post == nil {
		req, _ = http.NewRequest("GET", urls, nil)
	} else {
		postValues := url.Values{}
		for postKey, PostValue := range r.post{
			postValues.Set(postKey, PostValue)
		}
		postDataStr := postValues.Encode()
		postDataBytes := []byte(postDataStr)
		postBytesReader := bytes.NewReader(postDataBytes)
		req, _ = http.NewRequest("POST", urls, postBytesReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	//模拟浏览器访问
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	//配置host
	if r.host != "" {
		req.Host = r.host
	}
	resp,err := client.Do(req)
	//返回数据为空,请求不存在
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	//是否打印状态码
	if r.log {
		log.Println("[Status] "+strconv.Itoa(resp.StatusCode))
	}
	body,_ := ioutil.ReadAll(resp.Body)
	r.CurCookies = r.CurCookieJar.Cookies(req.URL)
	return string(body)
}

func (r *Request) PrintCookie() bool {
	length := len(r.CurCookies)
	flag := false
	for i :=0; i < length; i++ {
		var curCk *http.Cookie = r.CurCookies[i]
		if curCk.Value != "" {
			flag = true
		}
		log.Printf("[Cookie] %s\r\n",curCk.Value)
	}
	return flag
}