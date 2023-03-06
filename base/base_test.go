package base

import (
	"fmt"
	"testing"
)

func TestConvertCaseToCamel(t *testing.T) {

	n := ConvertCaseToCamel("a_b_c")
	fmt.Println("n", n)
}

func TestSortIntList(t *testing.T) {
	n := []int64{3, 5, 6, 1, 2}
	// n = SortIntList(n, false)
	fmt.Println("n", n)
}

func TestConvertMapToJson(t *testing.T) {
	mapData := make(map[string]interface{})
	mapData["age"] = 18
	mapData["name"] = "沉默小管"
	fmt.Println(ConvertMapToJson(mapData))
}

func TestConvertJsonToMap(t *testing.T) {
	mapData := "{\"age\":18,\"name\":\"沉默小管\"}"
	result := make(map[string]interface{})
	ConvertJsonToMap(mapData, &result)
	fmt.Println("res", result)
}

func TestConvertToString(t *testing.T) {
	n := int64(12132342)
	fmt.Println(ConvertToString(n, ""))
}
