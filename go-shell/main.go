package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	err := ShellOutput("ping www.baidu.com -c 5")
	if err != nil {
		fmt.Println("命令执行失败：", err)
	}
}

// ShellOutput shell实时输出, ScanLines以行输出。Words以单个词输出
func ShellOutput(shell string) error {
	cmd := exec.Command("/bin/bash", "-c", shell)
	reader, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println("shell执行出错：", err)
		return err
	}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("cmd.Wait()执行出错：", err)
		return err
	}
	return nil
}
