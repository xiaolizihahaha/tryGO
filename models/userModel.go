package models

import (
	"log"
	db "test1/database"
)

type User struct {
	User_id       int64  "json:user_id"
	User_name     string "json:user_name"
	User_password string "json:user_password"
}

func (u *User) AddUser() int64 {
	result, err := db.SqlDB.Exec("INSERT INTO user_test(username,password) VALUE(?,?)", u.User_name, u.User_password)

	if err != nil {
		log.Fatalln(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return id

}

func (u *User) Delete(id int64) int64 {
	result, err := db.SqlDB.Exec("DELETE FROM user_test WHERE id=?", id)

	if err != nil {
		log.Fatalln(err.Error())
	}

	ra, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return ra

}

func (u *User) Update() int64 {
	result, err := db.SqlDB.Prepare("UPDATE user_test SET username=?,password=? WHERE id=?")

	if err != nil {
		log.Fatalln(err.Error())
	}

	ra, err := result.Exec(u.User_name, u.User_password, u.User_id)

	if err != nil {
		log.Fatalln(err.Error())
	}

	sa, err := ra.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return sa
}

func (u *User) Select(id int64) User {
	err := db.SqlDB.QueryRow("SELECT id,username,password FROM user_test WHERE id=?", id).Scan(&u.User_id, &u.User_name, &u.User_password)
	if err != nil {
		log.Fatalln(err)
		u.User_id = 0
		u.User_name = "nil"
		u.User_password = "nil"

		return *u
	}
	return *u

}
