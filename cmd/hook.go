package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const hookScript = `#!/bin/sh
# gitpet post-commit hook — feeds your pet on every commit
gitpet feed 2>/dev/null
`

func runHook() error {
	if len(os.Args) < 3 {
		fmt.Println(`
  Usage:
    gitpet hook install    Install post-commit hook in current repo
    gitpet hook remove     Remove the hook
    gitpet hook status     Check if hook is installed
`)
		return nil
	}

	switch os.Args[2] {
	case "install":
		return installHook()
	case "remove", "uninstall":
		return removeHook()
	case "status":
		return hookStatus()
	default:
		return fmt.Errorf("unknown hook command: %s", os.Args[2])
	}
}

func getGitDir() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	if err != nil {
		return "", fmt.Errorf("not a git repository")
	}
	return strings.TrimSpace(string(out)), nil
}

func hookPath() (string, error) {
	gitDir, err := getGitDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(gitDir, "hooks", "post-commit"), nil
}

func installHook() error {
	path, err := hookPath()
	if err != nil {
		return err
	}

	// Check if hook already exists
	if data, err := os.ReadFile(path); err == nil {
		content := string(data)
		if strings.Contains(content, "gitpet") {
			fmt.Println("\n  ✅ gitpet hook is already installed!\n")
			return nil
		}
		// Append to existing hook
		newContent := content + "\n" + hookScript
		if err := os.WriteFile(path, []byte(newContent), 0755); err != nil {
			return fmt.Errorf("failed to update hook: %w", err)
		}
		fmt.Println("\n  ✅ gitpet hook appended to existing post-commit hook!\n")
		return nil
	}

	// Create hooks directory if needed
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create hooks dir: %w", err)
	}

	if err := os.WriteFile(path, []byte(hookScript), 0755); err != nil {
		return fmt.Errorf("failed to write hook: %w", err)
	}

	fmt.Printf("\n  ✅ gitpet hook installed at %s\n", path)
	fmt.Println("  Your pet will be fed on every commit! 🍖\n")
	return nil
}

func removeHook() error {
	path, err := hookPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("\n  No post-commit hook found.\n")
		return nil
	}

	content := string(data)
	if !strings.Contains(content, "gitpet") {
		fmt.Println("\n  No gitpet hook found in post-commit.\n")
		return nil
	}

	// Remove gitpet lines
	lines := strings.Split(content, "\n")
	var newLines []string
	skip := false
	for _, line := range lines {
		if strings.Contains(line, "gitpet post-commit hook") {
			skip = true
			continue
		}
		if skip && strings.Contains(line, "gitpet feed") {
			skip = false
			continue
		}
		skip = false
		newLines = append(newLines, line)
	}

	newContent := strings.TrimSpace(strings.Join(newLines, "\n"))

	if newContent == "#!/bin/sh" || newContent == "" {
		os.Remove(path)
		fmt.Println("\n  🗑️  gitpet hook removed!\n")
	} else {
		os.WriteFile(path, []byte(newContent+"\n"), 0755)
		fmt.Println("\n  🗑️  gitpet lines removed from post-commit hook!\n")
	}

	return nil
}

func hookStatus() error {
	path, err := hookPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("\n  ❌ No gitpet hook installed.\n  Run 'gitpet hook install' to set it up.\n")
		return nil
	}

	if strings.Contains(string(data), "gitpet") {
		fmt.Printf("\n  ✅ gitpet hook is active at %s\n\n", path)
	} else {
		fmt.Println("\n  ❌ No gitpet hook found in post-commit.\n  Run 'gitpet hook install' to set it up.\n")
	}
	return nil
}
