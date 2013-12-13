package gopenapi

import (
	"errors"
	"strconv"
)

type OpenApi struct {
	appid      int64
	appkey     string
	ServerName string
	format     string
}

func NewOpenApi(appid int64, appkey string) *OpenApi {
	return &OpenApi{appid: appid, appkey: appkey, ServerName: "openapi.tencentyun.com", format: "json"}
}

func (this *OpenApi) SetServerName(name string) {
	this.ServerName = name
}

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
