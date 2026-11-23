package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func fillPlaceholders(dir, name string) error {
	replacements := []struct {
		file, old, new string
	}{
		{"README.md", "# repo-name\n", "# " + name + "\n"},
		{"go.mod", "module github.com/" + org + "/repo-name\n", "module github.com/" + org + "/" + name + "\n"},
	}

	for _, r := range replacements {
		path := filepath.Join(dir, r.file)

		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		updated := strings.Replace(string(data), r.old, r.new, 1)
		if err := os.WriteFile(path, []byte(updated), info.Mode().Perm()); err != nil {
			return err
		}
	}

	return nil
}

func commitRepoName(dir string) error {
	steps := [][]string{
		{"add", "-A"},
		{"commit", "-m", "repo: fill in repo name placeholders"},
		{"push"},
	}

	for _, args := range steps {
		if err := git(dir, args...); err != nil {
			return fmt.Errorf("git %s: %w", args[0], err)
		}
	}

	return nil
}

func git(dir string, args ...string) error {
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
