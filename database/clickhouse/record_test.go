package clickhouse

import (
	"gitee.com/liqiyuworks/jf-go-kit/config"

	"gitee.com/liqiyuworks/jf-go-kit/base"

	"testing"
)

type TbTestUser struct {
	ID       int    `bson:"_id" json:"id" default:"62f0a1977bf6d11a444d6c33"`
	Username string `json:"username,omitempty" default:"jiufang"`
	Password string `json:"password,omitempty" default:"123456"`
}

func InitRecord() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitCkManager()
}

func TestInsertRecord(t *testing.T) {
	defer InitRecord()()
	var user = TbTestUser{Username: "ls", Password: "lsls12"}
	affected, err := InsertRecord("default", "t_user", &user)
	base.Glog.Infoln(">>> affected ", affected)

	if err != nil {
		t.Errorf("TestInsertRecord func test error, err = %v", err)
	}
}

func TestFindRecordsByCond(t *testing.T) {
	defer InitRecord()()
	var user []TbTestUser
	cond := "username = 'lisheng' or username = 'liqiyu'"
	err := FindRecordsByCond("default", "t_user", cond, 0, 10, &user)
	base.Glog.Infoln(">>> user ", user)

	if err != nil {
		t.Errorf("ExistMgoUserByUsername func test error, err = %v", err)
	}
}
