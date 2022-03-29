package repository

import (
	models "SQLite_JWT/model/user"
)

type UserRepo interface {
	FindUserByEmail(email string) (models.User, error)
	CheckLoginInfo(email string, password string) (models.User, error)
	Insert(u models.User) error
}
