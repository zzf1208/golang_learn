package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func Transaction(Db *sqlx.DB) {

	//开启事务
	tx, err := Db.Begin() //tx就是事务这次用来操作数据库的1个连接
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	_, err = tx.Exec("insert into user(name, age)values(?, ?)", "user0101", 108) //插入一条数据
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("update user set name=?, age=?", "user0101", 108)
	if err != nil {
		tx.Rollback() //如果有错误，则回滚
		return
	}

	err = tx.Commit() //提交事务
	if err != nil {
		tx.Rollback()
		return
	}
}

func main() {
	dns := "root:123456@tcp(127.0.0.1:3306)/golang"
	Db, err := sqlx.Connect("mysql", dns) //sqlx.connect连接数据库
	if err != nil {
		fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("ping failed, err:%v\n", err)
		return
	}

	fmt.Printf("connect to db succ\n")
	Transaction(Db)
}
