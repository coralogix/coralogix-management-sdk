//go:build linux

package main

import "syscall"

const prSetDumpable = 4

// A child running as the same user must not read this process's environment
// through /proc or inspect its memory. The Collector needs only the data key.
func protectProcessSecrets() error {
	_, _, errno := syscall.RawSyscall(syscall.SYS_PRCTL, prSetDumpable, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}
