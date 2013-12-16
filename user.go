package gopenapi

import (
	"encoding/json"
	"errors"
	"log"
)

type User struct {
	ApiAdapter
}

func NewUser(openid, openkey, pf string) *User {
	u := &User{}
	u.openid = openid
	u.openkey = openkey
	u.pf = pf
	return u
}

/**
 * 获取登录用户的信息，包括昵称、头像、性别等信息
 * @see http://wiki.open.qq.com/wiki/v3/user/get_info
 */
func (this *User) GetInfo(p map[string]string) (T_UserInfo, error) {
	var data T_UserInfo
	ret, err := this.api.Api("/v3/user/get_info", this.makeParams(p), "post", "http")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		}
		return data, nil
	}
}

/**
 * 获取登录用户的VIP信息
 * @see http://wiki.open.qq.com/wiki/v3/user/total_vip_info
 */
func (this *User) TotalVipInfo(p map[string]string) (T_UserVipInfo, error) {
	var data T_UserVipInfo
	ret, err := this.api.Api("/v3/user/total_vip_info", this.makeParams(p), "post", "http")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		}
		return data, nil
	}
}

/**
 * 验证登录用户是否黄钻，是否年费黄钻，如果是则返回其黄钻等级等信息。
 * @see http://wiki.open.qq.com/wiki/v3/user/is_vip
 */
func (this *User) IsVip() (T_UserIsVip, error) {
	var data T_UserIsVip
	ret, err := this.api.Api("/v3/user/is_vip", this.makeParams(nil), "post", "http")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		}
		return data, nil
	}
}

/**
 * 验证登录用户是否安装了应用
 * @see http://wiki.open.qq.com/wiki/v3/user/is_setup
 */
func (this *User) IsSetup() (bool, error) {
	ret, err := this.api.Api("/v3/user/is_setup", this.makeParams(nil), "post", "http")
	if err != nil {
		log.Fatal(err)
		return false, err
	} else {
		data := make(map[string]interface{})
		json.Unmarshal(ret, &data)

		if data["ret"] == 0.0 {
			return data["setuped"] == 1.0, nil
		} else {
			return false, errors.New("error")
		}
	}
}

/**
 * 验证用户的登录态，判断openkey是否过期，没有过期则对openkey有效期进行续期（一次调用续期2小时）。
 * @see http://wiki.open.qq.com/wiki/v3/user/is_login
 */
func (this *User) IsLogin() (bool, error) {
	ret, err := this.api.Api("/v3/user/is_login", this.makeParams(nil), "post", "http")
	if err != nil {
		log.Fatal(err)
		return false, err
	} else {
		var data T_RetBase
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return false, errors.New(data.Msg)
		} else {
			return true, nil
		}
	}
}

/**
 * 本接口仅适用于多区多服应用，用来验证用户登录态（即验证openkey），以及验证用户是否从选区页面（即验证seqid）进入应用。
 * 多区多服应用中设置有验证用户是否从选区页面进入应用的逻辑，将有助于防止用户直接通过修改应用地址的方式进入应用。
 * @see http://wiki.open.qq.com/wiki/v3/user/is_area_login
 */
func (this *User) IsAreaLogin(seqid string) (bool, error) {
	p := make(map[string]string)
	p["seqid"] = seqid
	ret, err := this.api.Api("/v3/user/is_area_login", this.makeParams(p), "post", "http")
	if err != nil {
		log.Fatal(err)
		return false, err
	} else {
		var data T_RetBase
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return false, errors.New(data.Msg)
		} else {
			return true, nil
		}
	}
}
