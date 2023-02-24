/*
 * @Author: lisheng
 * @Date: 2023-01-16 16:38:45
 * @LastEditTime: 2023-02-24 14:45:17
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /jf-go-kit/base/sort.go
 */
package base

import (
	"sort"

	"golang.org/x/exp/constraints"
)

/**
 * @description: 切片排序（支持多种类型船体）
 * @return {*}
 * @author: liqiyuWorks
 */
func SortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
