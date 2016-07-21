package global

import (
	"errors"
)

var (
	// common
	FileNotExists = errors.New("file not exists")
	FileReadFail  = errors.New("file read fail")

	// conf
	NotExistKey = errors.New("there is no this key")
	NotPointer  = errors.New("value is not a pointer")
	UnknowError = errors.New("other error")

	// sql
	RecordEmpty = errors.New("the condition get empty result")
	RecordError = errors.New("do sql comes out error")
	InsertError = errors.New("insert data comes out error")

	// cart
	CartEmpty = errors.New("cart is empty")
)
