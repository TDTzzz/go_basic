package dao

import (
	"testing"
)

func TestUserDAOImpl_Save(t *testing.T) {
	userDAO := &UserDAOImpl{}

	err := InitMysql("127.0.0.1", "3306", "root", "123456", "test_db")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user := &UserEntity{
		Username: "tdt",
		Password: "123456",
		Email:    "tdt@gmail.com",
	}
	err = userDAO.Save(user)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("new User ID is %d", user.ID)
}

func TestUserDAOImpl_SelectByName(t *testing.T) {
	userDAO := &UserDAOImpl{}

	err := InitMysql("127.0.0.1", "3306", "root", "123456", "test_db")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user, err := userDAO.SelectByName("tdt@gmail.com")
	t.Logf("result username is %s", user.Username)
}
