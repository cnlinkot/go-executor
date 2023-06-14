package rlog

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var loggerMap = make(map[string]*Logger)
var format = "2006-01-02"

// Tick 计算下一个0点的时间并且关闭所有logger
func Tick() {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, time.Local)
		t := time.NewTimer(next.Sub(now))
		<-t.C
		for _, e := range loggerMap {
			e.close = true
		}
	}
}

type Logger struct {
	name   string
	reader *bufio.Reader
	fs     *os.File
	close  bool
	//size 待定，已经写出的数据大小
	size int64
}

func New(name string, r io.ReadCloser) {
	logger := &Logger{name: name, reader: bufio.NewReader(r)}
	loggerMap[name] = logger
	go logger.rol()
}

func (l *Logger) rol() {
	for {
		bytes, err := l.reader.ReadBytes('\n')
		if err != nil {
			//子进程停止 pipe 关闭
			_ = l.fs.Close()
			delete(loggerMap, l.name)
			return
		}
		l.append(bytes)
	}
}

func (l *Logger) append(buf []byte) {
	if l.fs == nil {
		l.fs = open(l.name)
	}
	if l.close {
		err := l.fs.Close()
		if err != nil {
			println("无法关闭日志文件:", err.Error())
			return
		}
		l.fs = open(l.name)
		l.close = false
	}
	_, err := l.fs.Write(buf)
	if err != nil {
		println("无法写出日志:", err.Error())
	}
}

func open(name string) *os.File {
	_ = os.Mkdir("logs", 0750)
	fileName := genLogFileName(name)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		println("无法装载日志文件:", err.Error())
		return nil
	}
	return file
}

func genLogFileName(name string) string {
	return fmt.Sprintf("logs/%s-%s.log", name, time.Now().Format(format))
}
