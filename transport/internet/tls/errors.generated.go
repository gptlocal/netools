// Code generated by errors.codegen. DO NOT EDIT.
package tls
import "github.com/gptlocal/netools/common/errors"

type errPathObjHolder struct{}

func newError(values ...interface{}) *errors.Error {
	return errors.New(values...).WithPathObj(errPathObjHolder{})
}
