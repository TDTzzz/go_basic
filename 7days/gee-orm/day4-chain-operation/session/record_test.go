package session

import "testing"

var (
	user1 = &User{8, "Tom", 22}
	user2 = &User{9, "Jack", 23}
	user3 = &User{10, "Mike", 24}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	//err1 := s.DropTable()
	//err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

//func TestSession_Insert(t *testing.T) {
//	s := testRecordInit(t)
//	affected, err := s.Insert(user3)
//	if err != nil || affected != 1 {
//		t.Fatal("failed to create record")
//	}
//}
