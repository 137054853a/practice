package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "d://test1.txt"
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		fmt.Println("程序结束")
		return
	}
	//defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\r')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("程序结束")

	defer func() {
		file.Close()
		fmt.Print("关闭文件")
	}()
}
