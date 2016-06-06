package wallpaper

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

func Set(name string) error {
	result := win.SystemParametersInfo(20, 0, unsafe.Pointer(syscall.StringToUTF16Ptr(name)), 0)

	if !result {
		return errors.New("Failed to set background.")
	}
	return nil
}
