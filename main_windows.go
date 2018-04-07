// +build windows

package main

import (
	"syscall"

	sequences "github.com/konsorten/go-windows-terminal-sequences"
)

func initTerminal() {
	sequences.EnableVirtualTerminalProcessing(syscall.Stdout, true)
	sequences.EnableVirtualTerminalProcessing(syscall.Stderr, true)
}
