// 支付相关API
// Copyright 2013 lisijie

package gopenapi

import (
	"encoding/json"
	"errors"
	"log"
)

type Pay struct {
	ApiBase
}

func NewPay(openid, openkey, pf string) *Pay {
	p := &Pay{}
	p.openid = openid
	p.openkey = openkey
	p.pf = pf
	return p
}

/**
 * 用户点击按钮选择“购买”后，应用调用该接口发送请求，以获取本次交易的token，以及购买物品的url参数。
 * @see http://wiki.open.qq.com/wiki/v3/pay/buy_goods
 */
func (this *Pay) BuyGoods(params map[string]string) (T_BuyGoods, error) {
	var data T_BuyGoods
	ret, err := this.api.Api("/v3/pay/buy_goods", this.makeParams(params), "post", "https")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		} else {
			return data, nil
		}
	}
}

/**
 * 应用发货通知
 * @see http://wiki.open.qq.com/wiki/v3/pay/confirm_delivery
 */
func (this *Pay) ConfirmDelivery(params map[string]string) (T_RetBase, error) {
	var data T_RetBase
	ret, err := this.api.Api("/v3/pay/confirm_delivery", this.makeParams(params), "post", "https")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		} else {
			return data, nil
		}
	}
}

/**
 * 获取道具交易token
 * @see http://wiki.open.qq.com/wiki/v3/pay/exchange_goods
 */
func (this *Pay) ExchangeToken(params map[string]string) (T_BuyGoods, error) {
	var data T_BuyGoods
	ret, err := this.api.Api("/v3/pay/exchange_goods", this.makeParams(params), "post", "https")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		} else {
			return data, nil
		}
	}
}

/**
 * 获取用户游戏币余额
 * @see http://wiki.open.qq.com/wiki/v3/pay/get_balance
 */
func (this *Pay) GetBalance(params map[string]string) (T_Balance, error) {
	var data T_Balance
	ret, err := this.api.Api("/v3/pay/get_balance", this.makeParams(params), "post", "https")
	if err != nil {
		log.Fatal(err)
		return data, err
	} else {
		json.Unmarshal(ret, &data)
		if data.Ret > 0 {
			return data, errors.New(data.Msg)
		} else {
			return data, nil
		}
	}
}
