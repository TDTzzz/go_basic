package service

import (
	"context"
	"go_basic/go_kit_project/user/dao"
	"go_basic/go_kit_project/user/redis"
	"testing"
)

func TestUserServiceImpl_Login(t *testing.T) {
	err := dao.InitMysql("127.0.0.1", "3306", "root", "123456", "test_db")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = redis.InitRedis("127.0.0.1", "6379", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	userService := &UserServiceImpl{
		userDAO: &dao.UserDAOImpl{},
	}

	user, err := userService.Login(context.Background(), "tdt@mail.com", "123456")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("user id is %d", user.ID)

}

func TestUserServiceImpl_Register(t *testing.T) {

	err := dao.InitMysql("127.0.0.1", "3306", "root", "123456", "test_db")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = redis.InitRedis("127.0.0.1", "6379", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	userService := &UserServiceImpl{
		userDAO: &dao.UserDAOImpl{},
	}

	user, err := userService.Register(context.Background(),
		&RegisterUserVO{
			Username: "tt",
			Password: "123456",
			Email:    "tt@gmail.com",
		})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("user id is %d", user.ID)

}
