package point

import "errors"

/**
 * @description: 二分查找坐标对应的索引
 * @param {[]float64} dataSet: 数据集
 * @param {float64} findNum 需要查找的经纬度
 * @param {int} midIndex 中间索引 传0
 * @param {string} lOrRight 数据集是降序还是升序 desc or asc
 * @return {*}
 * @author: liqiyuWorks
 */
func BinarySearch(dataSet []float64, findNum float64, midIndex int, lOrRight string) ([]int, error) {
	indexList := []int{}

	if lOrRight == "asc" {
		if findNum > dataSet[len(dataSet)-1] {
			return indexList, errors.New("查找值超出阈值")
		}

		if findNum < dataSet[0] {
			return indexList, errors.New("查找值超出阈值")
		}
	} else {
		if findNum < dataSet[len(dataSet)-1] {
			return indexList, errors.New("查找值超出阈值")
		}
	}

	if len(dataSet) > 1 {
		mid := len(dataSet) / 2

		if dataSet[mid-1] <= findNum && findNum <= dataSet[mid] {
			indexList = append(indexList, midIndex+mid-1)
			indexList = append(indexList, midIndex+mid)
			return indexList, nil
		}

		if dataSet[mid-1] >= findNum && findNum >= dataSet[mid] {
			indexList = append(indexList, midIndex+mid-1)
			indexList = append(indexList, midIndex+mid)
			return indexList, nil
		}

		if dataSet[mid] <= findNum && findNum <= dataSet[mid+1] {
			indexList = append(indexList, midIndex+mid)
			indexList = append(indexList, midIndex+mid+1)
			return indexList, nil
		}

		if dataSet[mid+1] <= findNum && findNum <= dataSet[mid] {
			indexList = append(indexList, midIndex+mid)
			indexList = append(indexList, midIndex+mid+1)
			return indexList, nil
		}

		if findNum > dataSet[mid] {
			if lOrRight == "desc" {
				return BinarySearch(dataSet[:mid], findNum, midIndex, lOrRight)
			} else {
				midIndex += mid
				return BinarySearch(dataSet[mid+1:], findNum, midIndex+1, lOrRight)
			}
		} else {
			if lOrRight == "desc" {
				midIndex += mid
				return BinarySearch(dataSet[mid+1:], findNum, midIndex+1, lOrRight)
			} else {
				return BinarySearch(dataSet[:mid], findNum, midIndex, lOrRight)
			}
		}

	} else {
		return indexList, errors.New("dataSet's length must > 1")
	}
}
