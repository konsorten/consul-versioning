package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func runFromHash(rep *git.Repository, rev *plumbing.Reference) error {
	log.Infof("Target revision: %v", rev)

	// retrieve all commits
	commits, err := rep.Log(&git.LogOptions{
		From: rev.Hash(),
	})
	if err != nil {
		return fmt.Errorf("Failed to retrieve revision log from %v: %v", rev, err)
	}

	// handle commit changes
	var allChanges []gitChange
	var currentChange *gitChange

	for commit, err := commits.Next(); err == nil; commit, err = commits.Next() {
		// no patch for the latest entry
		if currentChange == nil {
			currentChange = &gitChange{Commit: commit}
			continue
		}

		// build patch
		patch, err := commit.Patch(currentChange.Commit)
		if err != nil {
			return fmt.Errorf("Failed to retrieve work-tree for revision %v: %v", commit.Hash, err)
		}

		currentChange.FilePatches = patch.FilePatches()

		// prepare for next iteration
		allChanges = append(allChanges, *currentChange)
		currentChange = &gitChange{Commit: commit}
	}

	return handleChanges(allChanges)
}
