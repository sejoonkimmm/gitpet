package cmd

import (
	"fmt"
	"os"
)

const version = "0.1.0"

const helpText = `
  🐾 gitpet v%s — A Tamagotchi that lives in your git history

  Usage:
    gitpet <command>

  Commands:
    init       Adopt a new pet
    status     Check on your pet (alias: s)
    feed       Feed your pet manually
    hook       Install/remove git hook
    graveyard  Visit fallen pets (alias: rip)
    version    Show version

  Your pet gets fed automatically on every git commit!
  Just run 'gitpet hook install' in your repo.

`

func Execute() error {
	if len(os.Args) < 2 {
		return runStatus()
	}

	switch os.Args[1] {
	case "init":
		return runInit()
	case "status", "s":
		return runStatus()
	case "feed":
		return runFeed()
	case "hook":
		return runHook()
	case "graveyard", "rip":
		return runGraveyard()
	case "version", "v", "--version", "-v":
		fmt.Printf("gitpet v%s\n", version)
		return nil
	case "help", "--help", "-h":
		fmt.Printf(helpText, version)
		return nil
	default:
		fmt.Printf(helpText, version)
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}
