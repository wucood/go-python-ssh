package main

import (
	"github.com/gin-gonic/gin"
	wt "github.com/wuhaomiao/tools"
)

func digitAppH5(c *gin.Context)  {
	// 实时执行结果
	wt.NewCMD("python -u /server/scripts/git-pull-digit-app-h5.py root 192.168.10.175 22")
	//cmd := exec.Command("/bin/bash", "-c", "python3 -u /Users/cat/golang/auto-git/git-pull-digit-app-h5.py root 192.168.10.175 22")
	//cmd := exec.Command("python3", "-u", "/Users/cat/golang/auto-git/git-pull-digit-app-h5.py", "root", "192.168.10.175", "22")

	// 查看详细错误
	//cmd := exec.Command("python", "-u", "/server/scripts/git-pull-digit-app-h5.py", "root", "192.168.10.175", "22")
	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	//	return
	//}
	//fmt.Println("Result111: " + out.String())
}