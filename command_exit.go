package main

import "os"

func commandExitCli(args ...string) error {
	os.Exit(3)
	return nil
}
