package dbconfiggenerator

import (
	"fmt"
	"os"
)

func GenDB() {
	fmt.Print("请输入数据库类型")
	var dbtype string
	fmt.Scanln(&dbtype)
	switch dbtype {
	case "mysql":
		mysqlconf()
	case "postgres":
		pgconf()
	case "sql server":
		sqlserverconf()
	case "tidb":
		mysqlconf()
	default:
		fmt.Println("未知数据库")
		os.Exit(3)
	}
}
