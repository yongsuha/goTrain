package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 替换为你的数据库连接信息
	dsn := "root:Dhrsny678bjsfb@!@tcp(117.78.2.33:3306)/TrainTicket?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("连接数据库失败: %v\n", err)
		return
	}
	defer db.Close()

	// 尝试 Ping 数据库
	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping 数据库失败: %v\n", err)
		return
	}
	fmt.Println("成功连接到数据库！")
}
