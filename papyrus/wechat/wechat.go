package wechat

import (
	"encoding/xml"
	"reflect"

	"github.com/kafrax/chaos"
	"github.com/kafrax/netask"
	"github.com/17bixin/gobase/papyrus/unite"
)

//todo Detecting the error of return system to avoid repeated payment

type HostList = string

const (
	ApiUnifyOrder HostList = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	ApiTransfer   HostList = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
)

type weChat struct {
	AppId     string `json:"app_id"`
	MchId     string `json:"mch_id"`
	AppKey    string `json:"-"`
	AppSecret string `json:"-"`
}

func newWeChat(pap *unite.Papyrus)*weChat{
	return &weChat{
		AppId:pap.AppID,
		MchId:pap.MchID,
		AppKey:pap.AppKey,
		AppSecret:pap.AppSecret,
	}
}

func (c *weChat) SignAndMarshal(m WePayer) ([]byte, error) {
	m.SetExtraParam()
	m.SetActionParam()
	data := m.GetMapData()
	e := reflect.ValueOf(c)
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") != "-" {
			data[t.Field(i).Tag.Get("json")] = e.Field(i).Interface().(string)
		}
	}
	//sign
	data["sign"] = genSign(data, c.AppKey)
	return xml.MarshalIndent(chaos.ForXmlMap(data), "", " ")
}

func (c *weChat) doPost(m WePayer) ([]byte, error) {
	b, err := c.SignAndMarshal(m)
	if err != nil {
		return nil, err
	}
	return netask.Post(
		m.ChooseHost(),
		"application/xml;charset=utf-8",
		false,
		b,
	)
}

func (c *weChat) CheckSign() error {
	return nil
}
