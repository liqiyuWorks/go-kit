package tdengine

import (
	"testing"
	"time"

	"gitee.com/liqiyuworks/jf-go-kit/config"
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
	exeSql := "select lon,lat,* from db_mfwam.t_mfwam_3651_1314 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) union all select lon,lat,* from db_mfwam.t_mfwam_3651_1315 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) union all select lon,lat,* from db_mfwam.t_mfwam_3652_1314 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) union all select lon,lat,* from db_mfwam.t_mfwam_3652_1315 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) ;"
	var obj []*TDMfwam
	RestFindRecordsPtr(exeSql, &obj)
	for _, v := range obj {
		t.Logf("time=%v, lat=%v, lon=%v, Seawavedirection=%v", v.Ts, v.Lat, v.Lon, v.Seawavedirection)
	}
}

func TestRestQuerySmocPtr(t *testing.T) {
	defer InitTdengine()()
	exeSql := "select lon,lat,* from db_smoc.t_smoc_3651_1314 where ts=1661990400000"
	var obj []*TDSmoc
	RestFindRecordsPtr(exeSql, &obj)
	for _, v := range obj {
		t.Logf("time=%v, lat=%v, lon=%v, seawateru=%v", v.Ts, v.Lat, v.Lon, v.Seawateru)
	}
}
