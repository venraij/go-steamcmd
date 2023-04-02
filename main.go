package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	windowsos "github.com/venraij/go-steamcmd/Windows"
	"golang.org/x/sys/windows"
)

func main() {
	Install()
}

func Install() {
	// Check if steamcmd is installed
	_, err := exec.LookPath("steamcmd")
	if err != nil {
		// Install SteamCMD
		if runtime.GOOS == "windows" {
			windowsos.Install()
		}
	}

	if !amAdmin() {
		runMeElevated()
	}
	time.Sleep(3 * time.Second)

	// Run SteamCMD
	log.Println("Running SteamCMD...")
	cmd := exec.Command("steamcmd", "+quit")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		log.Println("admin no")
		return false
	}
	log.Println("admin yes")
	return true
}
