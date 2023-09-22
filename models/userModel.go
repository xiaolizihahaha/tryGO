package models

import (
	"fmt"
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

func (u *User) SelectOne(id int64) User {
	rows, err := db.SqlDB.Query("select * from user_test where id=?", id)

	if err == nil {
		var u User
		for rows.Next() {
			err := rows.Scan(&u.User_id, &u.User_name, &u.User_password)
			if err == nil {
				return u
			}
			fmt.Println(err)

		}
	}
	fmt.Println(err)
	return User{User_id: 0, User_name: "none", User_password: "none"}
}

func SelectAll() []User {

	//先查询user表数据总数
	total := 0
	totals, err := db.SqlDB.Query("select count(*) from user_test")

	for totals.Next() {
		err := totals.Scan(&total)
		if err != nil {
			fmt.Println(err)
		}
	}

	Users := make([]User, 0, total)
	rows, err := db.SqlDB.Query("select * from user_test limit ?", total)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var u User
		err := rows.Scan(&u.User_id, &u.User_name, &u.User_password)
		if err != nil {
			fmt.Println(err)
		}
		Users = append(Users, u)

	}
	return Users
}

func SaveUsers(users []User) bool {
	fmt.Println(users)

	tx, _ := db.SqlDB.Begin()

	saveSign := true

	for _, u := range users {
		_, err := db.SqlDB.Exec("INSERT INTO user_test(username,password) VALUE(?,?)", u.User_name, u.User_password)

		if err != nil {
			saveSign = false
			fmt.Println(err)
		}

	}

	if saveSign == true {
		tx.Rollback()
		return saveSign
	}

	tx.Commit()
	return saveSign
}
