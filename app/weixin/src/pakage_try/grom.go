package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "root:123456@/bl_main?charset=utf8&parseTime=True&loc=Local")

	// add := Address{
	// 	MemberId:  10,
	// 	TrueName:  "123",
	// 	Gender:    1,
	// 	LiveArea:  "XXX",
	// 	Address:   "D$%$",
	// 	Mobile:    "18511280986",
	// 	IsDefault: 1,
	// }
	rest := db.Exec("UPDATE address SET member_id=member_id-(?) WHERE id = (?)", 1, 2)
	fmt.Printf("%#v\n", rest)
	fmt.Println(rest.Error)
	fmt.Println(rest.RowsAffected)
}

type Address struct {
	Id        int32 `gorm:"primary_key"`
	MemberId  int32
	TrueName  string
	Gender    int8
	LiveArea  string
	Address   string
	Mobile    string
	IsDefault int8
}

func (this *Address) TableName() string {
	return "address"
}
