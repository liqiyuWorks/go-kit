package base

import (
	"fmt"
	"testing"
)

func TestConvertCaseToCamel(t *testing.T) {
	n := ConvertCaseToCamel("a_b_c")
	fmt.Println("n", n)
}
