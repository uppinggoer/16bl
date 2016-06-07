package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "root:123456@/bl_main?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf("%v", db)
}
