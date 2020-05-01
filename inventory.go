package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type InventoryFile struct {
	// Active is the list of active machines
	Active Group `json:"active"`
	Meta   Meta  `json:"_meta"`
}

type Group struct {
	Hosts []string          `json:"hosts"`
	Vars  map[string]string `json:"vars"`
}

type Meta struct {
	Hostvars map[string]map[string]string `json:"hostvars"`
}

func (i *InventoryFile) GenerateFromInventory(inv Inventory) ([]byte, error) {
	domains, _ := inv.ListDomains()
	active := Group{
		Hosts: make([]string, len(domains)),
		Vars:  make(map[string]string),
	}
	for i, d := range domains {
		active.Hosts[i] = d.GetIP()
	}
	i.Active = active
	file, err := json.Marshal(i)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert struct to json")
	}
	return file, nil
}
