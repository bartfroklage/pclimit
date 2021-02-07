package main

import (
	"os/exec"
    "syscall"
    "unsafe"

	"golang.org/x/sys/windows"
)

const (
    DESKTOP_SWITCHDESKTOP = 0x0100 // The access to the desktop
)

// get desktop locked status
func IsWindowsLocked() bool {

    // load user32.dll only once
    user32 := windows.NewLazySystemDLL("user32.dll")

    openDesktop := user32.NewProc("OpenDesktopW")
    closeDesktop := user32.NewProc("CloseDesktop")
    switchDesktop := user32.NewProc("SwitchDesktop")

    var lpdzDesktopPtr uintptr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Default"))) //string
    var dwFloatsPtr uintptr = 0                                                               //uint32
    var fInheritPtr uintptr = 0                                                               //boolean
    var dwDesiredAccessPtr uintptr = uintptr(DESKTOP_SWITCHDESKTOP)                           //uint32

    r1, _, _ := syscall.Syscall6(openDesktop.Addr(), 4, lpdzDesktopPtr, dwFloatsPtr, fInheritPtr, dwDesiredAccessPtr, 0, 0)
    if r1 == 0 {
        panic("get desktop locked status error")
    }

    res, _, _ := syscall.Syscall(switchDesktop.Addr(), 1, r1, 0, 0)
    // clean up
    syscall.Syscall(closeDesktop.Addr(), 1, r1, 0, 0)

    return res != 1
}

func LockWindows() {
	cmd := exec.Command("rundll32.exe", "user32.dll,LockWorkStation")
	cmd.Start()
}