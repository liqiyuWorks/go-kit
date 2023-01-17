/*
 * @Author: lisheng
 * @Date: 2023-01-16 16:21:31
 * @LastEditTime: 2023-01-17 15:19:31
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/base/error.go
 */
package base

func CheckErr(err error, msg string) {
	if err != nil {
		Glog.Errorf("CheckErr => errMsg: %v \n", err)
	}
}
