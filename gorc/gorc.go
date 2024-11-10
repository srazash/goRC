package gorc

import (
	"fmt"
	"os"
)

const appLogo = `░█▀▀▀░▄▀▀▄░█▀▀▄░█▀▄░ 
░█░▀▄░█░░█░█▄▄▀░█░░░ 
░▀▀▀▀░░▀▀░░▀░▀▀░▀▀▀░ `
const appName = "gorc"
const appVers = "0.0.1"

func Run() {
	AppInfo()

	if len(os.Args) == 1 || os.Args[1] == "client" {
		Client()
	} else if os.Args[1] == "server" {
		Server()
	} else if os.Args[1] == "help" {
		Help()
	} else if os.Args[1] == "version" {
		Version()
	} else {
		Other()
	}
}

func AppInfo() {
	fmt.Print(appLogo)
	fmt.Print("\n")
	ShortVersion()
	fmt.Print("\n")
}

func Server() {
	fmt.Print("gorc is running in server mode 🖥️\npress ctrl+c or ctrl+z to kill the process\n")
}

func Client() {
	fmt.Print("gorc is running in client mode 👨‍💻\n")
}

func Help() {
	fmt.Print("gorc cannot help you right now 🧙‍♂️\n")
}

func ShortVersion() {
	fmt.Print(appName, " v", appVers, "\n")
}

func Version() {
	fmt.Print(appName, " v", appVers, " (", os.Args[0], ")\n")
}

func Other() {
	fmt.Print(os.Args[1], ": gorc doesn't understand this command 🩲\n")
	fmt.Print("check the command you typed and try again, or use the 'help' command 🆘\n")
}
