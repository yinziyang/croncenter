package main

import (
	"croncenter/local"
	"croncenter/sshtools"
	"fmt"
	"strings"
	"syscall"

	"flag"

	"golang.org/x/crypto/ssh/terminal"
)

var Input_Username = flag.String("u", local.UserName, "输入登录用户名")
var Input_Hostname = flag.String("s", "", "输入登录服务器名称,多个服务器使用,分割 (不允许为空)")

func init() {
	flag.Parse()
}

func checkInput() (err error) {
	if *Input_Hostname == "" {
		err = fmt.Errorf("hostname为空, 使用 -h 参数查看帮助")
		return
	}
	return
}

func promptPasswd() string {
	fmt.Print("Enter Password: ")
	if bytePassword, err := terminal.ReadPassword(int(syscall.Stdin)); err != nil {
		panic(err)
	} else {
		return strings.TrimSpace(string(bytePassword))
	}
}

func main() {
	var err error
	if err = checkInput(); err != nil {
		fmt.Println(err.Error())
		return
	}
	passwd := promptPasswd()
	for _, hostName := range strings.Split(*Input_Hostname, ",") {
		if err = sshtools.CreateTrust(hostName, *Input_Username, passwd); err != nil {
			fmt.Printf("|%s|\t%s\t%s\n", hostName, *Input_Username, err.Error())
		} else {
			fmt.Printf("|%s|\t%s\t%s\n", hostName, *Input_Username, "success")
		}
	}
}
