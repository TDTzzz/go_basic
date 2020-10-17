package session

import (
	"testing"
)

type users struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = &users{"hhh", 23}
	user2 = &users{"TDTzzz", 22}
)

func testRecordInit(t *testing.T) *Session {
	s := NewSession().Model(&users{})
	return s
}

func TestSession_Find(t *testing.T) {
	testRecordInit(t)
	//var users []user
	//if err := s.Find(&users); err != nil || len(users) != 2 {
	//	t.Fatal("failed to query all")
	//}
}
