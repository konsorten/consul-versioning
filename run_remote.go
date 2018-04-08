package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func runRemoteRepository(url string) error {
	rep, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:               url,
		NoCheckout:        true,
		RecurseSubmodules: git.NoRecurseSubmodules,
		SingleBranch:      true,
	})

	rev, err := rep.Head()
	if err != nil {
		return fmt.Errorf("Failed to get HEAD revision: %v", err)
	}

	return runFromHash(rep, rev)
}
