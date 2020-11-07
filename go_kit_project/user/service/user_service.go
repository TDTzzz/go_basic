package service

import (
	"context"
	"errors"
	"go_basic/go_kit_project/user/dao"
	"go_basic/go_kit_project/user/redis"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUserVO struct {
	Username string
	Password string
	Email    string
}

var (
	ErrUserExisted = errors.New("user is existed")
	ErrPassword    = errors.New("email and password are not match")
	ErrRegistering = errors.New("email is registering")
)

type UserService interface {
	//登录接口
	Login(ctx context.Context, email, password string) (*UserInfoDTO, error)
	//注册接口
	Register(ctx context.Context, vo *RegisterUserVO) (*UserInfoDTO, error)
}

type UserServiceImpl struct {
	userDAO dao.UserDAO
}

func MakeUserServiceImpl(userDAO dao.UserDAO) UserService {
	return &UserServiceImpl{userDAO: userDAO}
}

func (u UserServiceImpl) Login(ctx context.Context, email, password string) (*UserInfoDTO, error) {
	user, err := u.userDAO.SelectByEmail(email)
	if err == nil {
		if user.Password == password {
			return &UserInfoDTO{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			}, nil
		} else {
			return nil, ErrPassword
		}
	} else {
		log.Printf("err:%s", err)
	}
	return nil, err
}

func (u UserServiceImpl) Register(ctx context.Context, vo *RegisterUserVO) (*UserInfoDTO, error) {
	lock := redis.GetRedisLock(vo.Email, time.Duration(5)*time.Second)
	err := lock.Lock()
	if err != nil {
		log.Printf("err:%s", err)
		return nil, ErrRegistering
	}
	defer lock.Unlock()

	existUser, err := u.userDAO.SelectByEmail(vo.Email)

	if (err == nil && existUser == nil) || err == gorm.ErrRecordNotFound {
		newUser := &dao.UserEntity{
			Username: vo.Username,
			Password: vo.Password,
			Email:    vo.Email,
		}
		err = u.userDAO.Save(newUser)
		if err == nil {
			return &UserInfoDTO{
				ID:       newUser.ID,
				Username: newUser.Username,
				Email:    newUser.Email,
			}, nil
		}
	}

	if err == nil {
		err = ErrUserExisted
	}

	return nil, err
}
