package util

import (
	. "global"
	"reflect"

	"menteslibres.net/gosexy/to"
)

/*
	Returns the element of the Slice or Map given by route.
*/
func pick(src interface{}, dig bool, route ...interface{}) (*reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil, NotPointer
	}

	v = v.Elem()
	for _, key := range route {
		u := v
		switch v.Kind() {
		case reflect.Slice:
			switch i := key.(type) {
			case int:
				if i < v.Len() {
					v = v.Index(i)
				} else {
					return nil, NotExistKey
				}
			}
		case reflect.Map:
			vkey := reflect.ValueOf(key)
			v = v.MapIndex(vkey)
			if dig == true && v.IsValid() == false {
				u.SetMapIndex(vkey, reflect.MakeMap(u.Type()))
				v = u.MapIndex(vkey)
			}
			if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
				v = v.Elem()
			}
		default:
			return nil, NotExistKey
		}
		if v.IsValid() == true {
			if v.CanInterface() == true {
				v = reflect.ValueOf(v.Interface())
			}
		}
	}

	return &v, nil
}

/*
	Starts with src (pointer to Slice or Map) tries to follow the given route,
	if the route is found it then tries to copy or convert the found node into
	the value pointed by dst.
*/
func Get(src, dst interface{}, route ...interface{}) error {
	if len(route) < 1 {
		return NotExistKey
	}

	// check param type
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Ptr || srcValue.IsNil() {
		return NotPointer
	}

	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr || dstValue.IsNil() {
		return UnknowError
	}
	dstValue.Elem().Set(reflect.Zero(dstValue.Elem().Type()))

	// copy from src to dst
	p, err := pick(src, false, route...)
	if err != nil {
		return err
	}
	if p.IsValid() == false {
		return NotExistKey
	}

	if dstValue.Elem().Type() != p.Type() {
		// Trying conversion
		if p.CanInterface() == true {
			t, err := to.Convert(p.Interface(), dstValue.Elem().Kind())
			if err == nil {
				tv := reflect.ValueOf(t)
				if dstValue.Elem().Type() == tv.Type() {
					p = &tv
				}
			}
		}
	}

	if dstValue.Elem().Type() == p.Type() || dstValue.Elem().Kind() == reflect.Interface {
		dstValue.Elem().Set(*p)
	} else {
		return UnknowError
	}

	// return dstValue.Elem().Interface(), nil
	return nil
}
