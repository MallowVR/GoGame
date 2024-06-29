package main

import (
	"os"
	"os/user"
	"strings"
)

func main() {
	LoadConfig()
	LoadSkills()
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	var localPlayer player

	var command []string // = "stats"
	if len(os.Args) <= 1 {
		command = append(command, "stats")
	} else {
		for i := 1; i < len(os.Args); i++ {
			command = append(command, strings.ToLower(os.Args[i]))
		}
	}

	localPlayer.loadPlayer(user.Name)

	commandHandler(&localPlayer, command)

	localPlayer.savePlayer()

}
