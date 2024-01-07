package dbconfiggenerator

import "fmt"

func sqlserverconf() string {
	var ip, port, user, pass, dbname string
	fmt.Print("请输入ip:")
	fmt.Scanln(&ip)
	fmt.Print("请输入端口:")
	fmt.Scanln(&port)
	fmt.Print("请输入用户:")
	fmt.Scanln(&user)
	fmt.Print("请输入密码:")
	fmt.Scanln(&pass)
	fmt.Print("请输入数据库名:")
	fmt.Scanln(&dbname)
	res := "sqlserver://" + user + ":" + pass + "@" + ip + ":" + port + "?database=" + dbname
	fmt.Println("合成结果为:" + res)
	return res
}
