package main

import "os"

func commandExitCli(_ ...string) error {
	os.Exit(3)
	return nil
}
