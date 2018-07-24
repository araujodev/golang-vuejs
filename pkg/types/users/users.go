package users

import (
	"github.com/araujodev/golang-vuejs/pkg/db"
)

type Users []User

type User db.Users

func (u *User) TableName() string {
	return "users"
}
