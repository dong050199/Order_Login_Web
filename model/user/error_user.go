package model

import "errors"

var (
	ERR_USER_NOT_FOUND  = errors.New("ERR_USER_NOT_FOUND")
	ERR_EMAIL_DUPLICATE = errors.New("ERR_DUPLICATE_EMAIL")
)
