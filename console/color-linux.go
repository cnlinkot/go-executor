//go:build linux

package con

const (
	ColorSuccess = "\033[32m"
	ColorWarn    = "\033[33m"
	ColorError   = "\033[31m"
	ColorNone    = "\033[0m"
)

func setColor(color string) {
	print(color)
}
