package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile() {
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Printf("文件读取失败,err:%v", err)
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for {
		line, err := read.ReadString('\n')
		if err == io.EOF {
			fmt.Println("无可读内容")
			return
		}
		if err != nil {
			fmt.Printf("读取文件出错,err:%s", err)
			return
		}
		fmt.Print(line)
	}
}

func writeFile() {
	file, err := os.OpenFile("./2.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%s", err)
	}
	file.WriteString("测试文件写入")
	file.Close()
}

func useBufio() {
	var s string
	read := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入内容")
	s, _ = read.ReadString('\n')
	fmt.Printf("你输入的内容是:%s", s)
}

func main() {
	useBufio()
}
