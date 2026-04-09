package pet

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func dataDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".gitpet")
}

func statePath() string {
	return filepath.Join(dataDir(), "state.json")
}

func EnsureDir() error {
	return os.MkdirAll(dataDir(), 0755)
}

func Load() (*State, error) {
	data, err := os.ReadFile(statePath())
	if err != nil {
		if os.IsNotExist(err) {
			return &State{}, nil
		}
		return nil, err
	}
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func Save(s *State) error {
	if err := EnsureDir(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(statePath(), data, 0644)
}

func Exists() bool {
	_, err := os.Stat(statePath())
	return err == nil
}
