package main

import (
	"fmt"
	"os"

	"github.com/gsamokovarov/jump/cli"
	_ "github.com/gsamokovarov/jump/cmd"
	"github.com/gsamokovarov/jump/config"
)

func main() {
	args := cli.ParseArgs(os.Args)

	config, err := config.SetupDefault(os.Getenv("JUMP_HOME"))
	if err != nil {
		panic(fmt.Sprintf("bug: %s", err.Error()))
	}

	command, err := cli.DispatchCommand(args, "--help")
	if err != nil {
		panic(fmt.Sprintf("bug: %s", err.Error()))
	}

	command.Action(args.Rest(), config)
}
