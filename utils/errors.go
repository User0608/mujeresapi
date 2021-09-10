package utils

import (
	"errors"
)

var (
	ErrDataBaseError        = errors.New("No se pudo completar la operacion")
	ErrUsuarioNoDisponible  = errors.New("El usuario ya esta en uso")
	ErrUsuarioNotExists     = errors.New("El usuario no existe")
	ErrNullEntity           = errors.New("Null Entity")
	ErrExistUsername        = errors.New("Exist username")
	ErrNoUsernameOrPassword = errors.New("Missing Username or Password")
	ErrInvalidData          = errors.New("Datos invalidos")
	ErrNothingFind          = errors.New("Registro no encontrados")
	ErrIdIsNeeded           = errors.New("El id es necesario")
	ErrIDInvalido           = errors.New("El id proporcionado no es valido")
)
