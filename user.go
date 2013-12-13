package gopenapi

import (
	"log"
)

type User struct {
	openid  string
	openkey string
	pf      string
	api     *OpenApi
}

func NewUser(openid, openkey, pf string) *User {
	return &User{openid: openid, openkey: openkey, pf: pf}
}

func (this *User) SetApi(api *OpenApi) {
	this.api = api
}

func (this *User) makeParams() map[string]string {
	params := make(map[string]string)
	params["openid"] = this.openid
	params["openkey"] = this.openkey
	params["pf"] = this.pf
	return params
}

func (this *User) GetInfo() {
	params := this.makeParams()
	ret, err := this.api.Api("/v3/user/get_info", params, "post", "http")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(ret))
	}
}
