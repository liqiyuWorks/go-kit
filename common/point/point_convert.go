package point

import (
	"fmt"

	"gitee.com/wuxi_jiufang/jf-go-kit/base"
)

type Point struct {
	Lon float64 `json:"lon,omitempty"`
	Lat float64 `json:"lat,omitempty"`
}

type PointIndex struct {
	LonIndex int     `json:"lonIndex,omitempty"`
	LatIndex int     `json:"latIndex,omitempty"`
	Lon      float64 `json:"lon,omitempty"`
	Lat      float64 `json:"lat,omitempty"`
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
	newLonPoints, newLatPoints := []float64{}, []float64{}
	for i := range lonPoints {
		newLonPoints = append(newLonPoints, base.ConvertToFloat64(lonPoints[i], 0))
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

/**
 * @description: ConverGfsPointToIndex
 * @param {*} lonPoints
 * @param {[]string} latPoints
 * @param {float64} lon
 * @param {float64} lat
 * @return {*}
 * @author: liqiyuWorks
 */
func ConverGfsPointToIndex(lonPoints, latPoints []string, lon float64, lat float64) ([]int, []int, error) {
	newLonPoints, newLatPoints := []float64{}, []float64{}
	for i := range lonPoints {
		newLonPoints = append(newLonPoints, base.ConvertToFloat64(lonPoints[i], 0))
	}
	for i := range latPoints {
		newLatPoints = append(newLatPoints, base.ConvertToFloat64(latPoints[i], 0))
	}
	lonIndexs, lonErr := BinarySearch(newLonPoints, lon, 0, "asc")
	latIndexs, latErr := BinarySearch(newLatPoints, lat, 0, "desc")
	if lonErr != nil || latErr != nil {
		base.Glog.Errorf("ConverMfwamPointToIndex occur error: lon= %v,lat= %v ", lonErr, latErr)
		return nil, nil, fmt.Errorf("%v+%v", lonErr, latErr)
	}
	return lonIndexs, latIndexs, nil
}
