package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	. "global"

	yaml "gopkg.in/yaml.v2"
	"menteslibres.net/gosexy/dig"
)

type ConfType struct {
	confCommon map[interface{}]interface{}
	conf       map[interface{}]interface{}
}

func Conf(confName string) ConfType {
	strConfCommon := COMMON_CONF_PATH + confName + ".yaml"
	strConf := CONF_PATH + confName + ".yaml"

	// 通用的配置
	confCommon := make(map[interface{}]interface{})
	if Exist(strConfCommon) {
		file, err := os.Open(strConfCommon)
		if err != nil {
			// log warning
		} else {
			data, err := ioutil.ReadAll(file)
			if err != nil {
				// log warning
			}
			yaml.Unmarshal(data, &confCommon)
		}
	}

	// app 覆盖的配置
	conf := make(map[interface{}]interface{})
	if Exist(strConf) {
		file, err := os.Open(strConf)
		if err != nil {
			// log warning
		} else {
			data, err := ioutil.ReadAll(file)
			if err != nil {
				// log warning
			}
			yaml.Unmarshal(data, &conf)
		}
	}
	return ConfType{confCommon: confCommon, conf: conf}
}
func (self ConfType) String() string {
	return fmt.Sprintf("conf:%v\nconfCommon%v", self.conf, self.confCommon)
}
func (self ConfType) MustValue(key string, defaultValue interface{}) interface{} {
	value, err := self.Value(key)
	if nil != err {
		return defaultValue
	} else {
		return value
	}
}

// 根据key查找value  key="A.B.C"  即 v=map["A"]["B"]["C"]
func (self ConfType) Value(key string) (interface{}, error) {
	// 查找键值是否存在
	var value interface{}
	var err error
	if reflect.TypeOf(key).Kind() == reflect.String {
		// 获取所有配置
		if 0 >= len(key) {
			value = self.conf
			goto success
		}

		path := key
		var pathList = []string{path}
		if strings.Contains(path, ".") == true {
			pathList = strings.Split(path, ".")
			// return dig.Set(&self.values, value, route...)
		}

		// 拼装 route
		route := make([]interface{}, len(pathList))
		for i, _ := range pathList {
			route[i] = pathList[i]
		}

		err = dig.Get(&self.conf, &value, route...)
		if nil == err {
			goto success
		}
		err = dig.Get(&self.confCommon, &value, route...)
		if nil == err {
			goto success
		}
		goto fail
	}
	err = NotExistKey
fail:
	return value, err
success:
	valueType := reflect.ValueOf(value).Kind()
	switch valueType {
	case reflect.Slice, reflect.Map:
		break
	default:
		value = fmt.Sprint(value)
	}
	return value, nil
}

type AppConfType struct {
	conf map[interface{}]interface{}
}

func AppData(confName string, data interface{}) error {
	strConf := APP_CONF_PATH + confName + ".yaml"

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Ptr || dataValue.IsNil() {
		return NotPointer
	}

	if Exist(strConf) {
		file, err := os.Open(strConf)
		if err != nil {
			return FileNotExists
		} else {
			confData, err := ioutil.ReadAll(file)
			if err != nil {
				return FileReadFail
			}

			yaml.Unmarshal(confData, data)
		}
	} else {
		return FileNotExists
	}

	return nil
}
func AppConf(confName string) (AppConfType, error) {
	strConf := APP_CONF_PATH + confName + ".yaml"

	conf := make(map[interface{}]interface{})
	if Exist(strConf) {
		file, err := os.Open(strConf)
		if err != nil {
			// log warning
		} else {
			data, err := ioutil.ReadAll(file)
			if err != nil {
				// log warning
			}
			yaml.Unmarshal(data, &conf)
		}
	} else {
		return AppConfType{conf: conf}, FileNotExists
	}

	return AppConfType{conf: conf}, nil
}

func (self AppConfType) MustValue(key string, defaultValue interface{}) interface{} {
	value, err := self.Value(key)
	if nil != err {
		return defaultValue
	} else {
		return value
	}
}

// 根据key查找value  key="A.B.C"  即 v=map["A"]["B"]["C"]
func (self AppConfType) Value(key string) (interface{}, error) {
	// 查找键值是否存在
	var value interface{}
	var err error
	if reflect.TypeOf(key).Kind() == reflect.String {
		// 获取所有配置
		if 0 >= len(key) {
			value = self.conf
			goto success
		}

		path := key
		var pathList = []string{path}
		if strings.Contains(path, ".") == true {
			pathList = strings.Split(path, ".")
			// return dig.Set(&self.values, value, route...)
		}

		// 拼装 route
		route := make([]interface{}, len(pathList))
		for i, _ := range pathList {
			route[i] = pathList[i]
		}

		err = dig.Get(&self.conf, &value, route...)
		if nil == err {
			goto success
		}
		goto fail
	}
	err = NotExistKey
fail:
	return value, err
success:
	valueType := reflect.ValueOf(value).Kind()
	switch valueType {
	case reflect.Slice, reflect.Map:
		break
	default:
		value = fmt.Sprint(value)
	}
	return value, nil
}
