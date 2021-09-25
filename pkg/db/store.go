package db

import "github.com/k2rth1k/qt/model"

type Store interface {
	UserManagement
}

type UserManagement interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserWithEmail(email string) (*model.User, error)
}
