//go:build windows

package win

import (
	"syscall"
	"unsafe"
)

func SetTitle(title string) {
	dll := syscall.NewLazyDLL("kernel32.dll")
	setTitle := dll.NewProc("SetConsoleTitleW")
	ptr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return
	}
	_, _, _ = setTitle.Call(uintptr(unsafe.Pointer(ptr)))
}
