package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//insert()
	//query()
	//update()
	remove()
}

func insert() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/yeoman?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`insert user (user_name,user_age,user_sex) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("tony88", 25, 2)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func query() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/yeoman?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT  * from user")
	checkErr(err)

	//普通demo
	/*for rows.Next() {
		var userId int
		var userName string
		var userAge int
		var userSex int
		rows.Columns()
		err = rows.Scan(&userId, &userName, &userAge, &userSex)

		checkErr(err)
		fmt.Println(userId)
		fmt.Println(userName)
		fmt.Println(userAge)
		fmt.Println(userSex)
	}*/

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}

		}
		fmt.Println(record)
	}
}

//更新数据
func update() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/yeoman?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`update user set user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)

	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(num)
}

//删除数据
func remove() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/yeoman?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`delete from user where user_id=?`)
	checkErr(err)

	res, err := stmt.Exec(2)
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(num)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
