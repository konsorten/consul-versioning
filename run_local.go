package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
)

func runLocalRepository(dir string) error {
	rep, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf("Failed to open local Git repository: %v: %v", dir, err)
	}

	rev, err := rep.Head()
	if err != nil {
		return fmt.Errorf("Failed to get HEAD revision: %v", err)
	}

	return runFromHash(rep, rev)
}
