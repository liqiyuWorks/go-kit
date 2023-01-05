package base

import (
	"fmt"
	"testing"
)

/**
 * @description:
 * @param {*testing.T} t
 * @return {*}
 * @author: liqiyuWorks
 */
func TestStringSubA2B(t *testing.T) {
	dataStr := "2022-12-29 12:00:00"
	year := StringSubA2B(dataStr, 0, 4, false)
	mon := StringSubA2B(dataStr, 5, 7, false)
	day := StringSubA2B(dataStr, 8, 10, false)
	hour := StringSubA2B(dataStr, 11, 13, false)
	fmt.Println(year, mon, day, hour)
}
