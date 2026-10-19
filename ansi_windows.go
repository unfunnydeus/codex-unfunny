package main

import (
	"os"
	"syscall"
	"unsafe"
)

// Enable virtual terminal processing on Windows consoles. If stdout is not
// a console or the console is too old, fall back to a plain banner.
func enableANSI() bool {
	kernel := syscall.NewLazyDLL("kernel32.dll")
	getMode := kernel.NewProc("GetConsoleMode")
	setMode := kernel.NewProc("SetConsoleMode")
	handle := os.Stdout.Fd()
	var mode uint32
	ok, _, _ := getMode.Call(handle, uintptr(unsafe.Pointer(&mode)))
	if ok == 0 {
		return false
	}
	const virtualTerminalProcessing = 0x0004
	ok, _, _ = setMode.Call(handle, uintptr(mode|virtualTerminalProcessing))
	return ok != 0
}
