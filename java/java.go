package java

import (
	con "go-executor/console"
	"go-executor/exception"
	"go-executor/rlog"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var pidMap = make(map[int]string)
var cmdMap = make(map[int]exec.Cmd)
var args string
var javaPath string
var group = sync.WaitGroup{}

func PidMap() map[int]string {
	return pidMap
}

func CmdMap() map[int]exec.Cmd {
	return cmdMap
}

func Init(a []string) error {
	args = strings.Join(a, " ")
	path, err := exec.LookPath("java")
	if err != nil {
		con.Error("无法找到java，请检查环境配置")
		return err
	}
	con.Success("找到程序: %s", path)
	javaPath = path
	return nil
}

func Run(name string) error {
	con.Info("启动jar包：%s", name)
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return exception.InvalidArgumentsError{Msg: name + " 不存在"}
	}
	cmd := exec.Command(javaPath, "-jar", name, args)
	//cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	rlog.New(name, out)
	err = cmd.Start()
	if err != nil {
		return err
	}
	con.Success("子进程启动，pid: %d", cmd.Process.Pid)
	cmdMap[cmd.Process.Pid] = *cmd
	pidMap[cmd.Process.Pid] = name
	go delIfDone(name, *cmd)
	group.Add(1)
	return nil
}

func Kill(pid int) error {
	jar := pidMap[pid]
	if jar == "" {
		return exception.InvalidArgumentsError{Msg: "找不到进程"}
	}
	cmd := cmdMap[pid]
	if cmd.Process == nil {
		return exception.InvalidArgumentsError{Msg: "无对应进程"}
	}
	err := cmd.Process.Kill()
	if err != nil {
		return err
	}
	con.Info("结束进程: %v %v", pid, jar)
	return nil
}

func delIfDone(jar string, cmd exec.Cmd) {
	_ = cmd.Wait()
	con.Warn("%s 结束", jar)
	delete(cmdMap, cmd.Process.Pid)
	delete(pidMap, cmd.Process.Pid)
	group.Done()
}

func Wait() {
	group.Wait()
}
