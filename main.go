/*
/*
@Time : 2019/12/26 3:43 下午
@Author : tianpeng.du
@File : main.go
@Software: GoLand
*/
package main

import (
	"github.com/du2016/scheduler-framework-test/pkg"
	"log"
	"os"
)

func main()  {
	cmd:=pkg.Register()
	if err:=cmd.Execute();err!=nil {
		log.Println(err)
		os.Exit(1)
	}
}
