package service

import (
	"errors"
	"strings"
)

// StringService 提供一些字符串操作
type StringServiceInt interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringService struct{}

func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")