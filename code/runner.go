package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

func main() {
	// 运行命令
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out

	// 比对
	pipe, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
	}
	// 写入
	io.WriteString(pipe, "23 12\n")

	// 运行
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// 输出
	println(out.String())
}
