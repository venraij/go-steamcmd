package windows

import "github.com/venraij/go-chocolatey"

func Install() {
	chocolatey.Install()
	chocolatey.InstallPackage("SteamCMD", "-y")
}
