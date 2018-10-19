package tcore

import (
	"fmt"
)

//TError export
type TError struct {
	mText string
}

//Error export
func (pOwn *TError) Error() string {
	return pOwn.mText
}

//TNewError export
func TNewError(aFormat string, aParms ...interface{}) error {
	return &TError{mText: fmt.Sprintf(aFormat, aParms...)}
}
