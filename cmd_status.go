package main

import (
	"fmt"

	"github.com/lxc/incus/client"
)

func cmdStatus(c incus.InstanceServer, args []string) error {
	// Load the simulation
	routers, err := importFromLXD(c)
	if err != nil {
		return err
	}

	fmt.Printf("Number of routers: %d\n", len(routers))

	return nil
}
