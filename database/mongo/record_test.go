package mongo

import (
	"jf-go-kit/config"

	"jf-go-kit/base"

	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitMgo() func() error {
	config.Initialize("../../config/app.json") // 配置文件
	return InitMgocli()
}

type TbTestMgoUser struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id" default:"62f0a1977bf6d11a444d6c33"`
	Username  string             `json:"username,omitempty" default:"jiufang"`
	AccessKey string             `bson:"access_key" json:"access_key,omitempty" default:"123456"`
	Password  string             `json:"password,omitempty" default:"123456"`
}

func TestExistByQueryMap(t *testing.T) {
	defer InitMgo()()
	var user *TbTestMgoUser
	userQuery := map[string]string{"username": "lisheng"}
	affected, err := ExistByQueryMap("", "users", &userQuery, &user)
	base.Glog.Infoln(">>> affected ", affected)

	if err != nil {
		t.Errorf("ExistMgoUserByUsername func test error, err = %v", err)
	}
}

func TestInsertRecord(t *testing.T) {
	defer InitMgo()()
	user_id := primitive.NewObjectID()
	user := TbTestMgoUser{ID: user_id, Username: "lisheng", Password: "765432"}
	affected, err := InsertRecord("", "users", &user)
	base.Glog.Infoln(">>> affected ", affected)

	if err != nil {
		t.Errorf("ExistMgoUserByUsername func test error, err = %v", affected)
	}
}

func TestGetRecord(t *testing.T) {
	defer InitMgo()()
	userQuery := map[string]interface{}{"username": "lisheng"}
	user := TbTestMgoUser{}
	err := QueryRecord("", "users", &userQuery, &user)
	if user.Username != "" {
		base.Glog.Infoln("\n>>> user ", user)
	}

	if err != nil && user.Username == "" {
		t.Errorf("TestGetRecord func test error, err = %v", err)
	}
}

func TestFindRecords(t *testing.T) {
	defer InitMgo()()
	userQuery := map[string]string{"username": "jiufang"}
	userResList := []TbTestMgoUser{}
	err := FindRecords("", "users", &userQuery, &userResList)
	base.Glog.Infoln("\n>>> userMap ", userResList)
	// base.Glog.Infoln(">>> username ", userResList[0]["username"])

	if err != nil {
		t.Errorf("ExistMgoUserByUsername func test error, err = %v", userResList)
	}
}

func TestUpateRecord(t *testing.T) {
	defer InitMgo()()
	queryMap := map[string]string{"username": "lisheng"}
	updateMap := map[string]map[string]string{"$set": {"password": "lsls12"}}
	affected, err := UpateRecord("", "users", &queryMap, &updateMap)
	base.Glog.Infoln(">>> affected ", affected)

	if err != nil {
		t.Errorf("ExistMgoUserByUsername func test error, err = %v", affected)
	}
}

func TestDeleteRecord(t *testing.T) {
	defer InitMgo()()
	userQuery := map[string]string{"username": "lisheng"}
	affected, err := DeleteRecord("", "users", &userQuery)
	base.Glog.Infoln(">> affected", affected)

	if err != nil {
		t.Errorf("TestGetRecord func test error, err = %v", err)
	}
}
