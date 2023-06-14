package command

import (
	con "go-executor/console"
	"go-executor/exception"
	"go-executor/java"
	"strconv"
	"strings"
)

var handlerList []handler
var usages []string

// 命令处理器
type handler interface {
	//Match 命令是否能被该处理器处理
	Match(string) bool
	//Handle 处理命令
	Handle([]string) error
	//Convert 将命令转换为参数
	Convert(string) ([]string, error)
	//Usage 使用方法
	Usage() string
}

func Handle(cmd string) error {
	for _, h := range handlerList {
		if h.Match(cmd) {
			convert, err := h.Convert(cmd)
			if err != nil {
				return err
			}
			return h.Handle(convert)
		}
	}
	return exception.NoMatchingHandlerError{Msg: "输入help查看帮助"}
}

func addHandler(handler handler) {
	handlerList = append(handlerList, handler)
	usages = append(usages, handler.Usage())
}

func Init() {
	con.Info("开始初始化所有命令...")
	addHandler(helpHandler{})
	addHandler(pidHandler{})
	addHandler(killHandler{})
	addHandler(runHandler{})
	con.Success("命令初始化完成")
}

func generalConvert(cmd string, prefix string) []string {
	return strings.Split(cmd[len(prefix):], " ")
}

func generalMatch(cmd string, prefix string) bool {
	return strings.HasPrefix(cmd, prefix)
}

// help 命令 Handler
type helpHandler struct {
}

func (h helpHandler) Match(s string) bool {
	return generalMatch(s, "help")
}

func (h helpHandler) Handle(i []string) error {
	for _, u := range usages {
		con.Info(u)
	}
	return nil
}

func (h helpHandler) Convert(s string) ([]string, error) {
	return nil, nil
}

func (h helpHandler) Usage() string {
	return "帮助文档："
}

// pid 命令 Handler
type pidHandler struct {
}

func (p pidHandler) Match(s string) bool {
	return generalMatch(s, "pid")
}

func (p pidHandler) Handle(i []string) error {
	pidMap := java.PidMap()
	for pid, n := range pidMap {
		con.Info("%v : %v", pid, n)
	}
	return nil
}

func (p pidHandler) Convert(s string) ([]string, error) {
	return nil, nil
}

func (p pidHandler) Usage() string {
	return "pid - 无需参数，查询当前所有管理的子进程"
}

// kill 命令 Handler
type killHandler struct {
}

func (k killHandler) Match(s string) bool {
	return generalMatch(s, "kill")
}

func (k killHandler) Handle(i []string) error {
	pid, err := strconv.Atoi(i[0])
	if err != nil {
		return err
	}
	return java.Kill(pid)
}

func (k killHandler) Convert(s string) ([]string, error) {
	if len(s) < len("kill ") {
		return nil, exception.InvalidArgumentsError{Msg: "请输入正确的pid"}
	}
	return generalConvert(s, "kill "), nil
}

func (k killHandler) Usage() string {
	return "kill pid - 结束一个子进程，只能结束由该容器管理的进程"
}

// run 命令 Handler
type runHandler struct {
}

func (r runHandler) Match(s string) bool {
	return generalMatch(s, "run ")
}

func (r runHandler) Handle(i []string) error {
	return java.Run(i[0])
}

func (r runHandler) Convert(s string) ([]string, error) {
	if len(s) < len("run ") {
		return nil, exception.InvalidArgumentsError{Msg: "请输入正确的jar包名"}
	}
	return generalConvert(s, "run "), nil
}

func (r runHandler) Usage() string {
	return "run xx.jar - 运行并添加管理一个jar包"
}
