package request

import (
	"fmt"
	"testing"
)

func TestHttpGET(t *testing.T) {
	resData := make(map[string]interface{}, 0)
	err := GET("https://api.caiyunapp.com/v2.6/xxxxxx/120.296989,31.892651/realtime", nil, nil, &resData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resData)
}

type ResBaseModel struct {
	Code int    `bson:"code" json:"code,omitempty"`
	Msg  string `bson:"msg" json:"msg,omitempty"`
}

type LoginData struct {
	Token    string `json:"token"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

type LoginRes struct {
	ResBaseModel
	Data LoginData `json:"data"`
}

func TestHttpPOST(t *testing.T) {
	data := map[string]string{
		"username": "xxx",
		"password": "xxxx",
	}
	resData := &LoginRes{}

	err := POST("http://121.36.28.59:21630/user/login", nil, data, resData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resData.Data.Token)
}
