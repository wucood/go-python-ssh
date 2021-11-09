package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	userName := os.Args[1]
	passwd := os.Args[2]
	host := os.Args[3]
	cmdStr := fmt.Sprintf("python3 -u /server/scripts/wu-test.py")
	shell(userName, passwd, host, cmdStr)
}

func shell(userName, passwd, host, cmdStr string) {
	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
		User:            userName,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		fmt.Println("连接ssh服务器失败：", err)
		return
	}
	session, _ := client.NewSession()
	defer session.Close()

	// 定义远程session的输出，返回io.Reader类型
	reader, _ := session.StdoutPipe()

	// 接收实时输出的内容
	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}()

	// shell命令
	//cmdStr := fmt.Sprintf("python3 -u /server/scripts/wu-test.py")
	// 执行命令
	err = session.Run(cmdStr)
	if err != nil {
		fmt.Println("命令执行失败：", err)
		return
	}
}
