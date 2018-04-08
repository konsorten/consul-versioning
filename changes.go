package main

import (
	"fmt"

	fdiff "gopkg.in/src-d/go-git.v4/plumbing/format/diff"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type gitChange struct {
	Commit      *object.Commit
	FilePatches []fdiff.FilePatch
}

func handleChanges(changes []gitChange) error {

	return fmt.Errorf("not implemented")
}
