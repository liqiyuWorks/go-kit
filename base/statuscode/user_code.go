/*
 * @Author: lisheng
 * @Date: 2022-10-20 11:04:47
 * @LastEditTime: 2022-11-03 00:48:41
 * @LastEditors: lisheng
 * @Description: 用户模块状态码
 * @FilePath: /gin-mvc/base/statuscode/user_code.go
 */

package statuscode

var (
	OK_LOGIN       = &Status{1, "登录成功"}
	ERROR_LOGIN    = &Status{-1001, "登录失败"}
	OK_REGISTER    = &Status{1, "注册成功"}
	ERROR_REGISTER = &Status{-1002, "注册失败"}
)
