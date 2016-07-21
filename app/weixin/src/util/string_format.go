package util

import (
	"fmt"
	"reflect"
	"strconv"
)

// 数字转换为字符串
func GetMoneyStr(cent int64) string {
	if 0 >= cent {
		return "0"
	}
	str := fmt.Sprintf("%.2f", float64(cent)/100)
	return str
}

// 数字转换为字符串
func Itoa(inter interface{}) string {
	switch v := inter.(type) {
	case int64:
		return strconv.FormatInt(v, 10)
	case int, int32, int16, int8:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint, uint32, uint16, uint8:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	default:
		return "0"
	}
}

// 字符串换为数字
func Atoi(value string, bitSize int, signed bool) interface{} {
	switch bitSize {
	case 64:
		if signed {
			i, err := strconv.ParseInt(value, 10, 64)
			if nil != err {
				return int64(0)
			} else {
				return i
			}
		} else {
			i, err := strconv.ParseUint(value, 10, 64)
			if nil != err {
				return uint64(0)
			} else {
				return i
			}
		}
	case 32:
		if signed {
			i, err := strconv.ParseInt(value, 10, 32)
			if nil != err {
				i = 0
			}
			return int32(i)
		} else {
			i, err := strconv.ParseUint(value, 10, 32)
			if nil != err {
				i = 0
			}
			return uint32(i)
		}
	case 16:
		if signed {
			i, err := strconv.ParseInt(value, 10, 16)
			if nil != err {
				i = 0
			}
			return int16(i)
		} else {
			i, err := strconv.ParseUint(value, 10, 16)
			if nil != err {
				i = 0
			}
			return uint16(i)
		}
	case 8:
		if signed {
			i, err := strconv.ParseInt(value, 10, 8)
			if nil != err {
				i = 0
			}
			return int8(i)
		} else {
			i, err := strconv.ParseUint(value, 10, 8)
			if nil != err {
				i = 0
			}
			return uint8(i)
		}
	default:
		return 0
	}
}

// 提取 interface{} 中的数字  主要是应对 json Unmarshal
// 目前只做了整形
func MustNum(inter interface{}, bitSize int, signed bool) interface{} {
	var tmpNum int64
	tmpNum = 0

	switch inter.(type) {
	case string:
		return Atoi(inter.(string), bitSize, signed)
	case int64:
		tmpNum = inter.(int64)
	default:
	}

	switch bitSize {
	case 64:
		if signed {
			return tmpNum
		} else {
			return uint64(tmpNum)
		}
	case 32:
		if signed {
			return int32(tmpNum)
		} else {
			return uint32(tmpNum)
		}
	case 16:
		if signed {
			return int16(tmpNum)
		} else {
			return uint16(tmpNum)
		}
	case 8:
		if signed {
			return int8(tmpNum)
		} else {
			return uint8(tmpNum)
		}
	default:
		return 0
	}
}
