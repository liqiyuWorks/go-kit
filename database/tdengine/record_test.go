package tdengine

import (
	"testing"

	"gitee.com/liqiyuworks/jf-go-kit/config"
	"gitee.com/liqiyuworks/jf-go-kit/table"
)

func InitTdengine() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitTdClient()
}

func TestRestQueryMfwamPtr(t *testing.T) {
	defer InitTdengine()()
	exeSql := "select lon,lat,* from db_mfwam.t_mfwam_3651_1314 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) union all select lon,lat,* from db_mfwam.t_mfwam_3651_1315 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) union all select lon,lat,* from db_mfwam.t_mfwam_3652_1314 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) union all select lon,lat,* from db_mfwam.t_mfwam_3652_1315 where ts in (1665360000000, 1665370800000, 1665381600000, 1665392400000, 1665403200000, 1665414000000, 1665424800000, 1665435600000, 1665446400000, 1665457200000, 1665468000000, 1665478800000, 1665489600000, 1665500400000, 1665511200000, 1665522000000) ;"
	var obj []*table.TDMfwam
	RestFindRecordsPtr(exeSql, &obj)
	for _, v := range obj {
		t.Logf("time=%v, lat=%v, lon=%v, Seawavedirection=%v", v.Ts, v.Lat, v.Lon, v.Seawavedirection)
	}
}

func TestRestQuerySmocPtr(t *testing.T) {
	defer InitTdengine()()
	exeSql := "select lon,lat,* from db_smoc.t_smoc_3651_1314 where ts=1661990400000"
	var obj []*table.TDSmoc
	RestFindRecordsPtr(exeSql, &obj)
	for _, v := range obj {
		t.Logf("time=%v, lat=%v, lon=%v, seawateru=%v", v.Ts, v.Lat, v.Lon, v.Seawateru)
	}
}
