package base

import (
	"fmt"
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
	var ts int64 = 1671264300000
	fmt.Println(GetSpecifiedNumToTs(ts, 3))
}
