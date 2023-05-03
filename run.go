package main

import (
	"fmt"
	"os"
)

var isL int = 1

func islive(path string) {
	//检测指定文件是否存在
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if err != nil {
			panic(err)
		} else {
			isL = 0
		}
	}
}

func main() {
	islive("C:\\Windows\\System32\\vm3dservice.exe")
	islive("C:\\Program Files\\VMware\\VMware Tools\\vmtoolsd.exe")
	if isL == 0 {
		println("当前环境为VM环境")
	} else {
		println("不是VM环境")
	}
	fmt.Scanln()
}
