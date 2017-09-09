package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

/**
 * sql.Open函数实际上是返回一个连接池对象，不是单个连接。在open的时候并没有去连接数据库，只有在执行query、exce方法的时候才会去实际连接数据库。
 * 在一个应用中同样的库连接只需要保存一个sql.Open之后的db对象就可以了，不需要多次open。
 */
var db *sql.DB

func main() {
	startHttpServer()
}

//开启web服务
func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//db对象初始化
//声明一个全局的db对象，并进行初始化
/*func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	//SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
	db.SetMaxOpenConns(2000)
	//SetMaxIdleConns用于设置闲置的连接数。
	db.SetMaxIdleConns(1000)
	db.Ping()
}*/

func pool(w http.ResponseWriter, r *http.Request) {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	rows, err := db.Query("select * from user limit 1")
	defer rows.Close()
	checkErr(err)

	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)

	for rows.Next() {
		//将数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
	fmt.Println(w, "finish")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
