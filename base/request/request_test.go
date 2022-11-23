package request

import (
	"fmt"
	"testing"
)

func TestHttpDo(t *testing.T) {
	params := map[string]string{
		"name": "yuyu",
	}
	res, err := HttpDo("https://test.pet-dbc.cn/api/user/info", "GET",
		params, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	//post
	// herders := map[string]string{}
	// type Data struct {
	// 	UserId      int64 `json:"user_id"`
	// 	OrderStatus int64 `json:"order_status"`
	// 	PageSizeNum int64 `json:"page_size_num"`
	// 	Current     int64 `json:"current"`
	// }
	// data := Data{16, 0, 10, 1}
	// body, _ := json.Marshal(data)
	res, err = HttpDo("https://api.caiyunapp.com/v2.6/TAkhjf8d1nlSlspN/120.296989,31.892651/realtime", "GET",
		nil, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
