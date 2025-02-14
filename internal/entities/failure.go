package entities

import (
	"strings"
)

type Failure interface {
	Error() string
}

type failure struct {
	Code string
}

func NewFailure(code string) Failure {
	return &failure{Code: code}
}

func (f *failure) Error() string {
	return strings.ReplaceAll(f.Code, "_", " ")
}
