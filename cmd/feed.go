package cmd

import (
	"fmt"

	"github.com/sejoonkimmm/gitpet/pet"
)

func runFeed() error {
	if !pet.Exists() {
		fmt.Println("\n  🥚 No pet found! Run 'gitpet init' to adopt one.\n")
		return nil
	}

	state, err := pet.Load()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	if state.CurrentPet == nil {
		fmt.Println("\n  🥚 No pet found! Run 'gitpet init' to adopt one.\n")
		return nil
	}

	pet.IncrementStreak(state.CurrentPet)
	msg := pet.Feed(state.CurrentPet)
	fmt.Println(msg)

	return pet.Save(state)
}
