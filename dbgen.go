package dbconfiggenerator

import (
	"fmt"
	"os"
)

func GenDB(confpath string) {
	fmt.Print("请输入数据库类型:")
	var dbtype string
	var dburl string
	fmt.Scanln(&dbtype)
	switch dbtype {
	case "mysql":
		dburl = mysqlconf()
	case "postgres":
		dburl = pgconf()
	case "sqlserver":
		dburl = sqlserverconf()
	case "tidb":
		dburl = mysqlconf()
	default:
		fmt.Println("未知数据库")
		os.Exit(3)
	}
	strs := "dbtype : " + dbtype + "\ndburl : " + dburl
	f, err := os.Create(confpath)

	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(strs))
	}

}
