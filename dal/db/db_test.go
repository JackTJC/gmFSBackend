package db

import (
	"testing"

	"github.com/JackTJC/gmFS_backend/model"
)

func TestCreateUser(t *testing.T) {
	InitDB()
	err := User.CreateUser(&model.UserInfo{
		UserName: "test",
		Password: "passwd",
		Email:    "email",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("success")
}
