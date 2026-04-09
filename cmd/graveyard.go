package cmd

import (
	"fmt"

	"github.com/sejoonkimmm/gitpet/pet"
)

func runGraveyard() error {
	if !pet.Exists() {
		fmt.Println("\n  🌿 No gitpet data found. The graveyard is empty.\n")
		return nil
	}

	state, err := pet.Load()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	fmt.Println(pet.GraveyardText(state.Graveyard))
	return nil
}
