package operator

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库信息
const (
	USERNAME = "root"
	PASSWORD = "123482"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "play"
)

//定义user_tabel
type User struct {
	Id       int64  `json:"id"  form:"id"`
	UserName string `json:"username" form:"username"`
	Age      int16  `json:"age" form:"age"`
	Sex      string `json:"sex" form:"sex"`
	Location string `json:"location" form:"location"`
}

func Start() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	fmt.Println(connStr)
	DB, err := sql.Open("mysql", connStr)

	if err != nil {
		return nil, err
	}
	return DB, nil
}

func InsertUser(DB *sql.DB, user *User) (*sql.Result, error) {
	result, err := DB.Exec("insert INTO user_table(id, user_name, age, sex, location) values(?,?,?,?,?)",
		user.Id, user.UserName, user.Age, user.Sex, user.Location)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func QueryUsers(DB *sql.DB) (*[]User, error) {
	users := []User{}
	rows, err := DB.Query("select id,user_name,location from user_table")
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.UserName, &user.Location)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil

}
