//go:build windows

package con

//#include "windows.h"
import "C"

var handle = C.GetStdHandle(C.STD_OUTPUT_HANDLE)

const (
	ColorSuccess = C.FOREGROUND_GREEN | C.FOREGROUND_INTENSITY
	ColorWarn    = C.FOREGROUND_GREEN | C.FOREGROUND_RED | C.FOREGROUND_INTENSITY
	ColorError   = C.FOREGROUND_RED | C.FOREGROUND_INTENSITY
	ColorNone    = C.FOREGROUND_BLUE | C.FOREGROUND_GREEN | C.FOREGROUND_RED
)

func setColor(color int) {
	C.SetConsoleTextAttribute(handle, C.ushort(color))
}
