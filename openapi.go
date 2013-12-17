package gopenapi

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type OpenApi struct {
	appid      int64
	appkey     string
	ServerName string
	format     string
}

//实例化OpenAPI对象
func NewOpenApi(appid int64, appkey string) *OpenApi {
	return &OpenApi{appid: appid, appkey: appkey, ServerName: "openapi.tencentyun.com", format: "json"}
}

//设置OpenAPI服务器名
func (this *OpenApi) SetServerName(name string) {
	this.ServerName = name
}

//请求OpenAPI
func (this *OpenApi) Api(urlpath string, params map[string]string, method string, protocol string) ([]byte, error) {
	if _, ok := params["openid"]; !ok {
		return nil, errors.New("openkey is empty")
	}

	params["appid"] = strconv.FormatInt(this.appid, 10)
	params["format"] = this.format
	params["sig"] = MakeSign(method, urlpath, params, this.appkey+"&")

	url := protocol + "://" + this.ServerName + urlpath

	return request(url, params, method, protocol)
}

/**
 * 发送HTTP请求并获取返回内容
 *
 * @param string rqurl 请求URL
 * @param map[string]string params 请求参数,为键值都为string的map
 * @param string method 请求方式:get|post
 * @param string protocol 协议: http|https
 */
func request(rqurl string, params map[string]string, method string, protocol string) ([]byte, error) {
	var (
		err  error
		resp *http.Response
	)

	client := &http.Client{}

	if protocol == "https" {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	if method == "get" {
		querystring := makeQueryString(params)
		resp, err = client.Get(rqurl + "?" + querystring)
	} else {
		data := url.Values{}
		for k, v := range params {
			data.Set(k, v)
		}
		resp, err = client.PostForm(rqurl, data)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

//生成查询字符串
func makeQueryString(params map[string]string) string {
	str := make([]string, 0, len(params))
	for k, v := range params {
		str = append(str, url.QueryEscape(k)+"="+url.QueryEscape(v))
	}
	return strings.Join(str, "&")
}
