package gopenapi

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func request(rqurl string, params map[string]string, method string, protocol string) ([]byte, error) {
	var (
		err  error
		resp *http.Response
	)

	if method == "get" {
		querystring := makeQueryString(params)
		resp, err = http.Get(rqurl + "?" + querystring)
	} else {
		data := url.Values{}
		for k, v := range params {
			data.Set(k, v)
		}
		resp, err = http.PostForm(rqurl, data)
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

func makeQueryString(params map[string]string) string {
	str := make([]string, 0, len(params))
	for k, v := range params {
		str = append(str, url.QueryEscape(k)+"="+url.QueryEscape(v))
	}
	return strings.Join(str, "&")
}
