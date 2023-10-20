package main

import "fmt"

// list all the keys in my cache atm
func (cfg *config) commandDebugCache() error {
	for k, v := range cfg.pokeapiClient.C.C {
		fmt.Printf("key: %v, value: %+v\n", k, v)
	}
	return nil
}
