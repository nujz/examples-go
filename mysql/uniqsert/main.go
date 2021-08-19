/*

# 辅助表

CREATE TABLE `__uniqsert`(
	`id` INT UNSIGNED NOT NULL PRIMARY KEY
) ENGINE MYISAM CHARSET UTF8;

INSERT INTO `__uniqsert` VALUES (1);


# 业务表

CREATE TABLE `user`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE MYISAM CHARSET UTF8;

*/

package main

import (
	"database/sql"
	"fmt"

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
	// db.Close()

	st, err := db.Prepare("insert into `user`(`username`) select ? from `__uniqsert` where `id` = 1 and not exists (select `id` from `user` where `username` = ?)")
	checkError(err)

	username := "tom"
	res, err := st.Exec(username, username)
	checkError(err)

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}
