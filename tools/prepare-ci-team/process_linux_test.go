//go:build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestProtectProcessSecrets(t *testing.T) {
	if os.Getenv("CI_PROCESS_PROTECTION_HELPER") == "1" {
		parentEnv := fmt.Sprintf("/proc/%d/environ", os.Getpid())
		readEnv := func() error {
			command := exec.Command("cat", parentEnv)
			command.Env = []string{"PATH=/usr/bin:/bin"}
			return command.Run()
		}
		if err := readEnv(); err != nil {
			os.Exit(2)
		}
		if err := protectProcessSecrets(); err != nil {
			os.Exit(3)
		}
		if err := readEnv(); err == nil {
			os.Exit(4)
		}
		os.Exit(0)
	}

	command := exec.Command(os.Args[0], "-test.run=^TestProtectProcessSecrets$")
	command.Env = []string{
		"PATH=/usr/bin:/bin",
		"CI_PROCESS_PROTECTION_HELPER=1",
		"CI_FAKE_SECRET=local-test-only",
	}
	if err := command.Run(); err != nil {
		t.Fatal("a same-user child could read protected process credentials")
	}
}
