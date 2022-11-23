/*
 * @Author: lisheng
 * @Date: 2022-10-18 16:41:13
 * @LastEditTime: 2022-11-18 15:37:32
 * @LastEditors: lisheng
 * @Description: 表schema
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/table/table.go
 */
package table

import (
	"time"
)

type TbBaseModel struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
}

// TDengine - Mfwam表
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

// TDengine - Smoc表
type TDSmoc struct {
	Ts        time.Time `bson:"t'ts" json:"ts,omitempty"`
	Seawateru float64   `bson:"seawateru" json:"seawateru,omitempty"`
	Seawaterv float64   `bson:"seawaterv" json:"seawaterv,omitempty"`
	Lon       float64   `bson:"lon" json:"lon,omitempty"`
	Lat       float64   `bson:"lat" json:"lat,omitempty"`
}

// TDengine - Tuvssh表
type TDTuvssh struct {
	Ts                  time.Time `bson:"t'ts" json:"ts,omitempty"`
	Seawaterheight      float64   `bson:"seawaterheight" json:"seawaterheight,omitempty"`
	Seawatertemperature float64   `bson:"seawatertemperature" json:"seawatertemperature,omitempty"`
	Lon                 float64   `bson:"lon" json:"lon,omitempty"`
	Lat                 float64   `bson:"lat" json:"lat,omitempty"`
}

type TbUserInfo struct {
	TbBaseModel
	Username string `gorm:"type:varchar(128);index;"`
	Password string `gorm:"type:varchar(512);"`
}

//指定数据库表名称
func (TbUserInfo) TableName() string {
	return "t_user_info"
}

func InitRegisterTables() {
	// mysql.RegisterTable("", &TbUserInfo{})
	// pg.RegisterTable("", &TbUserInfo{})
}
