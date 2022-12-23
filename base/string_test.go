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
	str := "liqiyuworks"
	fmt.Println(StringSubA2B(str, 1, 3, true))
}
