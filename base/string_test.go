package base

import (
	"fmt"
	"testing"
)

func TestStringSubA2B(t *testing.T) {
	str := "liqiyuworks"
	fmt.Println(StringSubA2B(str, 1, 3, true))
}
