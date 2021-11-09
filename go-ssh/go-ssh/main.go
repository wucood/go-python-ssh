package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func main() {
	userName := "wucood"
	passwd := "wu123456"
	host := "192.168.10.163:22"

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
	//cmdStr := fmt.Sprintf("ping www.baidu.com -c 5")
	cmdStr := fmt.Sprintf("python3 -u /server/scripts/wu-test.py")
	// 执行命令
	err = session.Run(cmdStr)
	if err != nil {
		fmt.Println("命令执行失败：", err)
		return
	}
}
