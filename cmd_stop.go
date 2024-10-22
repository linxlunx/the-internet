package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/lxc/incus/client"
	"github.com/lxc/incus/shared/api"
)

func cmdStop(c incus.InstanceServer, args []string) error {
	var wgBatch sync.WaitGroup

	if os.Getuid() != 0 {
		return fmt.Errorf("Container stop must be run as root.")
	}

	// Load the simulation
	routersMap, err := importFromLXD(c)
	if err != nil {
		return err
	}

	routers := []*Router{}
	for _, v := range routersMap {
		if v.Tier < 1 || v.Tier > 3 {
			continue
		}

		routers = append(routers, v)
	}

	// Helper function
	stopContainer := func(name string) {
		defer wgBatch.Done()

		req := api.InstanceStatePut{
			Action:  "stop",
			Timeout: -1,
			Force:   true,
		}

		op, err := c.UpdateInstanceState(name, req, "")
		if err != nil {
			return
		}

		err = op.Wait()
		if err != nil {
			return
		}
	}

	// Stop the containers
	batch := 8
	batches := len(routers) / batch
	remainder := len(routers) % batch

	current := 0
	for i := 0; i < batches; i++ {
		for j := 0; j < batch; j++ {
			wgBatch.Add(1)
			go stopContainer(routers[current].Name)
			current += 1
		}
		wgBatch.Wait()
	}

	for k := 0; k < remainder; k++ {
		wgBatch.Add(1)
		go stopContainer(routers[current].Name)
		current += 1
	}
	wgBatch.Wait()

	// Destroy all the interfaces
	err = networkDestroy(routersMap)
	if err != nil {
		return err
	}

	return nil
}
