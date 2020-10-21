package session

import (
	"go_basic/7days/gee-orm/day5-hooks/log"
	"testing"
)

type Account struct {
	Id       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before insert", account)
	account.Id += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "111122"
	return nil
}

func TestSession_CallMethod(t *testing.T) {

}
