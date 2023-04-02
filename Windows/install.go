package windowsos

import (
	"log"

	"github.com/venraij/go-chocolatey"
)

func Install() {
	chocoExists := chocolatey.IsInstalled()
	if !chocoExists {
		chocolatey.Install()
	}

	log.Println("Installing SteamCMD...")
	chocolatey.InstallPackage("SteamCMD", "-y")
}
