package utils

import (
	"math/rand"
	"strings"
	"time"
)

func RandonID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

func ValidarDecimal(target string) error {
	nums := "0123456789"
	estado := 1
	for _, s := range target {
		switch estado {
		case 1:
			if s == '-' {
				estado = 2
			} else if strings.ContainsRune(nums, s) {
				estado = 3
			} else {
				return ErrInvalidData
			}
		case 2:
			if strings.ContainsRune(nums, s) {
				estado = 3
			} else {
				return ErrInvalidData
			}
		case 3:
			if s == '.' {
				estado = 4
			} else if !strings.ContainsRune(nums, s) {
				return ErrInvalidData
			}
		case 4:
			if strings.ContainsRune(nums, s) {
				estado = 5
			} else {
				return ErrInvalidData
			}
		case 5:
			if !strings.ContainsRune(nums, s) {
				return ErrInvalidData
			}
		}
	}
	if estado == 3 || estado == 5 {
		return nil
	}
	return ErrInvalidData
}
