package con

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

func out(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}

func Info(msg string, args ...any) {
	lock.TryLock()
	out(msg, args...)
	lock.Unlock()
}

func Warn(msg string, args ...any) {
	lock.TryLock()
	setColor(ColorWarn)
	out(msg, args...)
	setColor(ColorNone)
	lock.Unlock()
}

func Success(msg string, args ...any) {
	lock.TryLock()
	setColor(ColorSuccess)
	out(msg, args...)
	setColor(ColorNone)
	lock.Unlock()
}

func Error(msg string, args ...any) {
	lock.TryLock()
	setColor(ColorError)
	out(msg, args...)
	setColor(ColorNone)
	lock.Unlock()
}
