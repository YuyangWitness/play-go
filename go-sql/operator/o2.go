package operator

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User2 struct {
	Id       int64  `gorm:"NOT NULL;primaryKey;" json:"id"  form:"id"`
	Username string `gorm:"NOT NULL;" json:"username" form:"username"`
	Age      int16  `gorm:"NOT NULL;" json:"age" form:"age"`
	Sex      string `gorm:"NOT NULL;" json:"sex" form:"sex"`
	Location string `gorm:"NOT NULL;" json:"location" form:"location"`
}

func StartO2() (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	DB.SingularTable(true)

	return DB, nil
}

func ListUser(DB *gorm.DB) *[]User2 {
	var users []User2
	DB.Table("user_table").Find(&users)
	return &users
}

func CreateUser(DB *gorm.DB, user *User2) {
	DB.Table("user_table").Create(&user)
}
