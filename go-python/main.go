package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "-u", "/Users/cat/github/go-python-ssh/sshCMD.py", "root", "192.168.10.163", "22")
	reader, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println("命令执行失败：", err)
		return
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("cmd.Wait()命令执行失败：", err)
		return
	}
}
