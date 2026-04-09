package cmd

import (
	"fmt"

	"github.com/sejoonkimmm/gitpet/pet"
)

func runInit() error {
	state, err := pet.Load()
	if err != nil {
		return fmt.Errorf("failed to load state: %w", err)
	}

	// If there's a living pet, bury it first
	if state.CurrentPet != nil && state.CurrentPet.Alive {
		fmt.Printf("\n  ⚠️  You already have a pet: %s (%s)\n", state.CurrentPet.Name, state.CurrentPet.Stage.Korean())
		fmt.Printf("  Adopting a new pet will send %s to the graveyard.\n", state.CurrentPet.Name)
		fmt.Printf("  Run 'gitpet status' to see your current pet.\n\n")
		return nil
	}

	// Bury dead pet
	if state.CurrentPet != nil && !state.CurrentPet.Alive {
		grave := pet.KillPet(state.CurrentPet, "starvation")
		state.Graveyard = append(state.Graveyard, grave)
	}

	newPet := pet.NewPet()
	state.CurrentPet = newPet

	if err := pet.Save(state); err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}

	art := pet.GetArt(pet.Egg, pet.Happy)
	fmt.Printf(`
  🎉 Welcome to gitpet! 🎉

  You adopted a new pet!%s
  Name: %s
  Stage: %s

  Feed your pet by making git commits!
  Run 'gitpet hook install' to auto-feed on every commit.

`, art, newPet.Name, newPet.Stage.Korean())

	return nil
}
