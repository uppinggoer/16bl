package util

import (
	"reflect"
	"sort"
	"strings"
)

// 根据元素的特定字段排序数组
// @param inputList 待排序数组
// @param Key 用来排序的Key
// return []inertface{},error

// 辅助排序的结构体
type sortFactory struct {
	Asc   bool          // 升序
	Entry reflect.Value // Slice
	Key   string
}

func (self *sortFactory) Len() int {
	if self.Entry.Kind() != reflect.Slice {
		return 0
	}
	return self.Entry.Len()
}
func (self *sortFactory) Less(i, j int) bool {
	valueI := self.Entry.Index(i)
	if valueI.Kind() == reflect.Struct {
		// 判断有没有 Key
		valueI = valueI.FieldByName(self.Key)
	} else if valueI.Kind() == reflect.Map {
		valueI = valueI.MapIndex(reflect.ValueOf(self.Key))
	}

	// 获取 valueJ 信息
	valueJ := self.Entry.Index(j)
	if valueJ.Kind() == reflect.Struct {
		// 判断有没有 Key
		valueJ = valueJ.FieldByName(self.Key)
	} else if valueJ.Kind() == reflect.Map {
		valueJ = valueJ.MapIndex(reflect.ValueOf(self.Key))
	}

	// 字段无效，则不排序
	if !valueI.IsValid() || !valueJ.IsValid() || valueJ.Kind() != valueJ.Kind() {
		return (i < j) == self.Asc
	}

	// 进行排序
	switch valueI.Kind() {
	case reflect.String:
		return strings.Compare(valueI.String(), valueJ.String()) < 0 == self.Asc
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return valueI.Int() < valueJ.Int() == self.Asc
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return valueI.Uint() < valueJ.Uint() == self.Asc
	case reflect.Float32, reflect.Float64:
		return valueI.Float() < valueJ.Float() == self.Asc
	default:
		// 暂不支持
		return (i < j) == self.Asc
	}
}

func (self sortFactory) Swap(i, j int) {
	tmpI := reflect.ValueOf(self.Entry.Index(i).Interface())
	self.Entry.Index(i).Set(self.Entry.Index(j))
	self.Entry.Index(j).Set(tmpI)
}

func SortList(inputList interface{}, sortKey string, asc bool) interface{} {
	entryValue := reflect.ValueOf(inputList)
	if entryValue.Kind() != reflect.Slice {
		return inputList
	}

	sortTmp := &sortFactory{Entry: entryValue, Key: sortKey, Asc: asc}
	sort.Sort(sortTmp)

	return sortTmp.Entry.Interface()
}
