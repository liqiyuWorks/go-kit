/*
 * @Author: lisheng
 * @Date: 2023-02-24 14:40:31
 * @LastEditTime: 2023-02-24 17:23:17
 * @LastEditors: lisheng
 * @Description:
 * @FilePath: /go-kit/base/sort_test.go
 */
package base

import (
	"fmt"
	"testing"
)

func TestSortSlice(t *testing.T) {
	floatSlice := []float64{2.3, 1.2, 0.2, 51.2}
	SortSlice(floatSlice, true)
	fmt.Println(floatSlice)
}
