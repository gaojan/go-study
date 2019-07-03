package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接数据库 "user:password@tcp(host:port)/dbname?参数=参数值"
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/xrt_open")
	defer db.Close()
	if err != nil {
		fmt.Println("打开数据库失败：", err)
		return
	}
	// 检查连接是否成功
	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	fmt.Println("连接数据库成功")

	// 一、查询数据，Query方法查询，返回sql.Rows结果集
	rows, err := db.Query("select id, name from t_id where id=5")
	if err != nil {
		fmt.Println("查询数据失败：", err)
		return
	}
	var id int
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name)
	}

	// 二、插入数据, 通过Exec方法插入数据，返回的结果是一个sql.Result类型 , 切记这里的占位符是?
	rest, err := db.Exec("insert into t_haman(name,age) values (?,?)", "gaojian", 30)
	if err != nil {
		fmt.Println("插入数据失败：", err)
		return
	}
	// 通过LastInsertId可以获取插入数据的id
	userId, err := rest.LastInsertId()
	fmt.Println("id>>", userId)
	// 通过RowAffected可以获取受影响的行数
	rowCount, err := rest.RowsAffected()
	fmt.Println("受影响的行:", rowCount)

	// 三、更新数据
	result, err := db.Exec("update t_haman set name=? where id=?", "golang", 1)
	if err != nil {
		fmt.Println("更新数据失败:", err)
		return
	}
	fmt.Println("更新数据成功：")
	fmt.Println(result.RowsAffected())

	// 四、 删除数据
	result1, err := db.Exec("delete from t_haman where id=?", 2)
	if err != nil {
		fmt.Println("删除数据失败:", err)
	}
	fmt.Println("删除数据成功")
	fmt.Println(result1.RowsAffected())

	rows.Close()
}
