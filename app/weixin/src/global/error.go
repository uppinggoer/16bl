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
)
