package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	//var shell string = "python D:\\python\\ssh-shell\\test.py"
	cmd := exec.Command("python", "-u","D:\\python\\ssh-shell\\ssh-shell.py", "root","192.168.10.163","22")
	//cmd := exec.Command("python", "-u","D:\\python\\ssh-shell\\test.py")
	reader,_ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil{
		fmt.Println("命令执行失败：",err)
		return
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
	}
	err = cmd.Wait()
	if err != nil{
		fmt.Println("cmd.Wait()命令执行失败：",err)
		return
	}
}
