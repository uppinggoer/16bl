package global

import (
	"errors"
)

var (
	// common
	FileNotExists = errors.New("file not exists")
	FileReadFail  = errors.New("file read fail")

	// http
	HttpMethodNot = errors.New("http method error")

	// conf
	NotExistKey = errors.New("there is no this key")
	NotPointer  = errors.New("value is not a pointer")
	UnknowError = errors.New("other error")

	// sql
	RecordEmpty = errors.New("the condition get empty result")
	RecordError = errors.New("do sql comes out error")
	InsertError = errors.New("insert data comes out error")

	// key
	ReidsKeyEmpty    = errors.New("redis key cannot be nil")
	ReidsPasswdEmpty = errors.New("redis passwd cannot be nil")

	// cart
	CartEmpty = errors.New("cart is empty")
	// order
	NotYourOrder = errors.New("not your order")

	// wechat
	JsAuthError = errors.New("fail get js auth")

	// xml
	XmlParseEmpty = errors.New("xml empty")
)
