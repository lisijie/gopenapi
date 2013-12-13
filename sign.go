package gopenapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"log"
	"net/url"
	"sort"
	"strings"
)

/**
 * 生成签名
 */
func MakeSign(method string, urlpath string, params map[string]string, secret string) string {
	mk := makeSource(method, urlpath, params)
	sign := hmac_sha1(mk, []byte(secret))
	return base64encode([]byte(sign))
}

func makeSource(method string, urlpath string, params map[string]string) []byte {
	var buf bytes.Buffer
	buf.WriteString(strings.ToUpper(method))
	buf.WriteString("&")
	buf.WriteString(url.QueryEscape(urlpath))
	buf.WriteString("&")

	qs := make([]string, 0)
	p := ksort(params)
	for k, v := range p {
		qs = append(qs, k+"="+v)
	}

	buf.WriteString(url.QueryEscape(strings.Join(qs, "&")))

	return buf.Bytes()
}

func ksort(data map[string]string) map[string]string {
	keys := make([]string, 0, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	result := make(map[string]string)
	for _, v := range keys {
		result[v] = data[v]
	}

	return result
}

func base64encode(src []byte) string {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write(src)
	enc.Close()
	return buf.String()
}

func hmac_sha1(src, key []byte) []byte {
	h := hmac.New(sha1.New, key)
	if _, err := h.Write(src); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}
