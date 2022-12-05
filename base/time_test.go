package base

import (
	"testing"
)

func TestGetDate(t *testing.T) {
	date := GetDate()
	Glog.Infoln("date=", date)
	if date == "" {
		t.Errorf("TestGetDate error, date = %v", date)
	}
}

func TestGetLastHour(t *testing.T) {
	var ts int64 = 1667038333000
	lastTs := GetLastHourToTs(ts)
	if lastTs == 0 {
		t.Errorf("TestGetDate error, lastTs = %v", lastTs)
	}
}
