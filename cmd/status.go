package cmd

import (
	"fmt"

	"github.com/sejoonkimmm/gitpet/pet"
)

func runStatus() error {
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

	fmt.Println(pet.StatusText(state.CurrentPet))

	// Save updated state (hunger/happiness may have changed)
	return pet.Save(state)
}
