package session

import "testing"

type User struct {
	Id   int
	Name string
	Age  int
}

func (*User) TableName() string {
	return "users2"
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	//_ = s.DropTable()
	_ = s.CreateTable()
	//if !s.HasTable() {
	//	t.Fatal("Failed to create table User")
	//}
}

//func TestSession_Model(t *testing.T) {
//	s := NewSession().Model(&User{})
//	table := s.RefTable()
//	s.Model(&Session{})
//	if table.Name != "users" || s.RefTable().Name != "Session" {
//		t.Fatal("Failed to change model")
//	}
//}
