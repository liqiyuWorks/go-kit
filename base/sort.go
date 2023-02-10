package base

import "sort"

type SortIntType []interface{}

//元素个数
func (t SortIntType) Len() int {
	return len(t)
}

//比较结果
func (t SortIntType) Less(i, j int) bool {
	// return t[i] < t[j]
	return true
}

//交换方式
func (t SortIntType) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func SortIntList(sortList SortIntType, reverse bool) SortIntType {
	if reverse {
		sort.Sort(sort.Reverse(sortList))
	} else {
		sort.Sort(sortList)
	}
	return sortList
}
