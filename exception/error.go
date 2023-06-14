package exception

type NoMatchingHandlerError struct {
	Msg string
}

func (n NoMatchingHandlerError) Error() string {
	return "没有那样的命令: " + n.Msg
}

type InvalidArgumentsError struct {
	Msg string
}

func (i InvalidArgumentsError) Error() string {
	return "参数无效: " + i.Msg
}
