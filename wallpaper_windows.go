package wallpaper

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

func Set(name string) {
	win.SystemParametersInfo(20, 0, unsafe.Pointer(syscall.StringToUTF16Ptr(name)), 0)
}
