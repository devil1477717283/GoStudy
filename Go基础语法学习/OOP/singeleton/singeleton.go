package User

import (
	"sync"
)

type user struct {
	Id   int
	Name string
	Sex  string
	Age  int
}

func newUser() *user {
	return &user{
		Id:   0,
		Name: "张三",
		Sex:  "Dont' Know",
		Age:  -1,
	}
}

var (
	u        *user
	userOnce sync.Once
)

func GetUserInstance() *user {
	userOnce.Do(func() {
		if u == nil {
			u = newUser()
		}
	})
	return u
}
