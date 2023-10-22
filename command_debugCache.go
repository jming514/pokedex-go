package main

import "fmt"

// list all the keys in my cache atm
func (cfg *config) commandDebugCache(_ ...string) error {
	for k := range cfg.pokeapiClient.C.C {
		fmt.Printf("key: %+v\n", k)
	}
	return nil
}
