package point

import (
	"fmt"

	"gitee.com/liqiyuworks/jf-go-kit/base"
)

type Point struct {
	Lon float64 `json:"lon,omitempty"`
	Lat float64 `json:"lat,omitempty"`
}

type PointIndex struct {
	Ilon int `json:"lon,omitempty"`
	Ilat int `json:"lat,omitempty"`
}

/**
 * @description: ConverMfwamPointToIndex
 * @param {*} lonPoints
 * @param {[]string} latPoints
 * @param {float64} lon
 * @param {float64} lat
 * @return {*}
 * @author: liqiyuWorks
 */
func ConverMfwamPointToIndex(lonPoints, latPoints []string, lon float64, lat float64) ([]int, []int, error) {
	newLonPoints := []float64{}
	newLatPoints := []float64{}
	for i := range lonPoints {
		newLonPoints = append(newLonPoints, base.Decimal2(base.ConvertToFloat64(lonPoints[i], 0)))
	}
	for i := range latPoints {
		newLatPoints = append(newLatPoints, base.ConvertToFloat64(latPoints[i], 0))
	}
	lonIndexs, lonErr := BinarySearch(newLonPoints, lon, 0, "asc")
	latIndexs, latErr := BinarySearch(newLatPoints, lat, 0, "asc")
	if lonErr != nil || latErr != nil {
		base.Glog.Errorf("ConverMfwamPointToIndex occur error: lon= %v,lat= %v ", lonErr, latErr)
		return nil, nil, fmt.Errorf("%v+%v", lonErr, latErr)
	}
	return lonIndexs, latIndexs, nil
}
