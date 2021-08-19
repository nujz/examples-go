/*

CREATE TABLE `user`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NOT NULL,
	`created_at` DATE NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE MYISAM CHARSET UTF8;

*/

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "test@/test?charset=utf8")
	checkError(err)
	// defer db.Close()

	st, err := db.Prepare("INSERT INTO `user` SET `username`=?,created_at=?")
	checkError(err)
	r, err := st.Exec("tom", time.Now())
	checkError(err)
	id, err := r.LastInsertId()
	checkError(err)
	fmt.Println(id)

	st1, err := db.Prepare("SELECT `username`,`created_at` FROM `user` WHERE id = ?")
	checkError(err)
	rows, err := st1.Query("1 or 1")
	checkError(err)

	for rows.Next() {
		var username string
		var created_at string
		err = rows.Scan(&username, &created_at)
		checkError(err)

		fmt.Println(username, created_at)
	}

	st2, err := db.Prepare("UPDATE `user` SET `username` = ? WHERE `id` = ?")
	checkError(err)
	aff, err := st2.Exec("bob", id)
	checkError(err)
	fmt.Println(aff.RowsAffected())

	st3, err := db.Prepare("DELETE FROM `user` WHERE `id` = ?")
	checkError(err)
	aff1, err := st3.Exec(id)
	checkError(err)
	fmt.Println(aff1.RowsAffected())
}
