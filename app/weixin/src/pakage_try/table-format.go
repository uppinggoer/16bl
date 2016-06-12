package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var typeMap = map[string]string{
	"mediumint": "int32",
	"tinyint":   "int8",
	"int":       "int64",
	"varchar":   "string",
	"timestamp": "string",
}

func main() {
	file, err := os.Open("/tmp/table.info")
	if nil != err {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var mapRow = map[string]string{}
	var arrRow = []string{}
	for {
		line, err := reader.ReadBytes('\n')
		if nil != err {
			break
		}

		strLine := string(line)
		strLine = strings.Trim(strLine, " ")

		arrField := strings.Split(strLine, " ")
		if 0 >= len(arrField) {
			continue
		}

		var rowInfo string
		if '`' == arrField[0][0] {
			// field name
			stop := len(arrField[0])
			filedName := arrField[0][1 : stop-1]
			arrName := strings.Split(filedName, "_")

			arrName[0] = strings.ToUpper(arrName[0][0:1]) + arrName[0][1:]
			if 1 < len(arrName) {
				arrName[1] = strings.ToUpper(arrName[1][0:1]) + arrName[1][1:]
				arrName[0] += arrName[1]
			}
			// fmt.Println(arrName[0])

			// type
			arrField[1] = strings.Split(arrField[1], "(")[0]
			arrField[1] = typeMap[arrField[1]]
			// fmt.Println(arrField[1])

			rowInfo = arrName[0] + " " + arrField[1]
			mapRow[filedName] = rowInfo
			arrRow = append(arrRow, filedName)
		}
		if strings.HasPrefix(strings.ToUpper(strLine), "PRIMARY") {
			primaryKey := strings.Split(strLine, "`")[1]
			if _, ok := mapRow[primaryKey]; ok {
				mapRow[primaryKey] += "\t `gorm:\"primary_key\"`"
			}
		}
	}

	for _, field := range arrRow {
		fmt.Println(mapRow[field])
	}
}
