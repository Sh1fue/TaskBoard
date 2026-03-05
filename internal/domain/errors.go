package domain


import (
	"errors"
)


var (
	ErrUserNotFound = errors.New("Пользователь не найден")
	ErrUserExist = errors.New("Пользователь с таким email уже существует")
)