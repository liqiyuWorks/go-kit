package tdengine

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"gitee.com/liqiyuworks/go-kit/base"
	"gitee.com/liqiyuworks/go-kit/config"
)

type TDMfwam struct {
	Ts                 time.Time `bson:"t'ts" json:"ts,omitempty"`
	Seawavedirection   float64   `bson:"seawavedirection" json:"seawavedirection,omitempty"`
	Seawaveheight      float64   `bson:"seawaveheight" json:"seawaveheight,omitempty"`
	Seawaveperiod      float64   `bson:"seawaveperiod" json:"seawaveperiod,omitempty"`
	Swellwavedirection float64   `bson:"swellwavedirection" json:"swellwavedirection,omitempty"`
	Swellwaveheight    float64   `bson:"swellwaveheight" json:"swellwaveheight,omitempty"`
	Swellwaveperiod    float64   `bson:"swellwaveperiod" json:"swellwaveperiod,omitempty"`
	Windwavedirection  float64   `bson:"windwavedirection" json:"windwavedirection,omitempty"`
	Windwaveheight     float64   `bson:"windwaveheight" json:"windwaveheight,omitempty"`
	Windwaveperiod     float64   `bson:"windwaveperiod" json:"windwaveperiod,omitempty"`
	Lon                float64   `bson:"lon" json:"lon,omitempty"`
	Lat                float64   `bson:"lat" json:"lat,omitempty"`
}

type TDSmoc struct {
	Ts        time.Time `bson:"t'ts" json:"ts,omitempty"`
	Seawateru float64   `bson:"seawateru" json:"seawateru,omitempty"`
	Seawaterv float64   `bson:"seawaterv" json:"seawaterv,omitempty"`
	Lon       float64   `bson:"lon" json:"lon,omitempty"`
	Lat       float64   `bson:"lat" json:"lat,omitempty"`
}

func InitTdengine() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitTdClient()
}

func TestRestQueryMfwamPtr(t *testing.T) {
	defer InitTdengine()()
	exeSql := "select lon,lat,* from db_mfwam.t_mfwam_3651_1314 where ts = 1665360000000"
	var obj []*TDMfwam
	RestFindRecordsPtr("", exeSql, &obj)
	for _, v := range obj {
		t.Logf("time=%v, lat=%v, lon=%v, Seawavedirection=%v", v.Ts, v.Lat, v.Lon, v.Seawavedirection)
	}
}

func TestRestQuerySmocPtr(t *testing.T) {
	defer InitTdengine()()
	exeSql := "select lon,lat,* from db_smoc.t_smoc_3651_1314 where ts=1661990400000"
	var obj []*TDSmoc
	RestFindRecordsPtr("", exeSql, &obj)
	for _, v := range obj {
		t.Logf("time=%v, lat=%v, lon=%v, seawateru=%v", v.Ts, v.Lat, v.Lon, v.Seawateru)
	}
}

type TestIntList []int64

//元素个数
func (t TestIntList) Len() int {
	return len(t)
}

//比较结果
func (t TestIntList) Less(i, j int) bool {
	return t[i] < t[j]
}

//交换方式
func (t TestIntList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func TestRestQueryTDMaps(t *testing.T) {
	defer InitTdengine()()
	exeSql := "show databases"
	reMaps, _ := RestFindRecordsMap("", exeSql)
	// var mfwamForeDate, smocForeDate int64 = 0, 0
	var mfwamlist, smocList []int64
	for k := range reMaps {
		name := base.ConvertToString(reMaps[k]["name"], "")
		splitNameList := strings.Split(name, "_")
		if len(splitNameList) >= 4 {
			// 此时进行对比
			if splitNameList[1] == "mfwam" {
				newDate := base.ConvertToInt64(splitNameList[3], 0)
				mfwamlist = append(mfwamlist, newDate)

			} else if splitNameList[1] == "smoc" {
				newDate := base.ConvertToInt64(splitNameList[3], 0)
				smocList = append(smocList, newDate)
			}
		}
	}

	fmt.Printf("mfwamlist=%d, smocList=%d \n ", mfwamlist, smocList)
	sort.Sort(TestIntList(mfwamlist))
	sort.Sort(TestIntList(smocList))
	fmt.Printf("sorted => mfwamlist=%d, smocList=%d \n ", mfwamlist, smocList)
}
