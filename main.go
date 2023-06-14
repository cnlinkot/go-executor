package main

import (
	"bufio"
	"go-executor/command"
	con "go-executor/console"
	"go-executor/java"
	"go-executor/rlog"
	"go-executor/win"
	"os"
	"strings"
	"syscall"
)

func main() {
	con.Info("父进程启动：Pid -> %d", syscall.Getpid())
	args := os.Args[1:]
	err := java.Init(args)
	if err != nil {
		con.Error(err.Error())
		return
	}
	win.SetTitle("go-executor | sa@linkot.cn")
	command.Init()
	dir, err := os.ReadDir(".")
	if err != nil {
		con.Error("无法扫描jar包: %s", err.Error())
		return
	}
	for _, file := range dir {
		name := file.Name()
		if strings.HasSuffix(name, ".jar") {
			err := java.Run(name)
			if err != nil {
				con.Error("jar包启动失败: %s", err.Error())
			}
		}
	}
	go rlog.Tick()
	go handleCommand()
	java.Wait()
	con.Info("父进程退出")
}

func handleCommand() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		str, err := reader.ReadString('\n')
		if err != nil {
			con.Error("命令错误：%s", err.Error())
		}
		err = command.Handle(strings.Trim(str, "\n\r"))
		if err != nil {
			con.Error("命令执行错误：%s", err.Error())
		}
	}
}
