package utils

import (
	"errors"
)

var (
	ErrNullEntity           = errors.New("Null Entity")
	ErrExistUsername        = errors.New("Exist username")
	ErrNoUsernameOrPassword = errors.New("Missing Username or Password")
	ErrInvalidData          = errors.New("Datos invalidos")
	ErrNothingFind          = errors.New("Registro no encontrados")
	ErrIdIsNeeded           = errors.New("El id es necesario")
)
