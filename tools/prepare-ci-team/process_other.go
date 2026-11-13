//go:build !linux

package main

func protectProcessSecrets() error { return nil }
