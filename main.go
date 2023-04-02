package SteamCMD

import (
	"log"
	"os/exec"
	"runtime"

	windows "github.com/venraij/SteamCMDWrapper/Windows"
)

func Install() {
	if runtime.GOOS == "windows" {
		windows.Install()
	}

	cmd := exec.Command("steamcmd")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
