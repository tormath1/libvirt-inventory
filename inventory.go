package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {
	p, err := NewProvider(DEFAULT_CONNECT_URI)
	if err != nil {
		fmt.Printf("unable to create libvirt provider: %v", err)
		os.Exit(1)
	}
	domains, _ := p.ListDomains()
	active := Group{
		Hosts: make([]string, len(domains)),
		Vars:  make(map[string]string),
	}
	for i, d := range domains {
		active.Hosts[i] = d.GetIP()
	}
	i := InventoryFile{Active: active}

	inv, _ := json.Marshal(i)
	fmt.Println(string(inv))

}
