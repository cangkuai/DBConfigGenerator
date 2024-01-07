package dbconfiggenerator

import "fmt"

func pgconf() string {
	var ip, port, user, pass, dbname, timez string
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
	fmt.Print("请输入时区(默认为上海):")
	fmt.Scanln(&timez)
	if timez == "" {
		timez = "Asia/Shanghai"
	}
	res := "host=" + ip + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=" + timez
	fmt.Println("合成结果为:" + res)
	return res
}
