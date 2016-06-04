package wallpaper

import (
	"fmt"
	"syscall"
	"unsafe"
)

func Set(name string) {
	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("SystemParametersInfoW")

	var SPI_SETDESKWALLPAPER = 0x0014

	ret, _, _ := proc.Call(0,
		uintptr(SPI_SETDESKWALLPAPER),
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(0))
	fmt.Printf("Return: %d\n", ret)
}
