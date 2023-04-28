package common

//go:generate go run github.com/gptlocal/netools/common/errors/codegen

import "github.com/gptlocal/netools/common/errors"

// ErrNoClue is for the situation that existing information is not enough to make a decision. For example, Router may return this error when there is no suitable route.
var ErrNoClue = errors.New("not enough information for making a decision")

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must2 panics if the second parameter is not nil, otherwise returns the first parameter.
func Must2(v interface{}, err error) interface{} {
	Must(err)
	return v
}

// Error2 returns the err from the 2nd parameter.
func Error2(v interface{}, err error) error {
	return err
}
