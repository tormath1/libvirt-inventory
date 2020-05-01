package main

import (
	"fmt"
	"os"
)

func main() {
	var i InventoryFile
	p, err := NewProvider(DEFAULT_CONNECT_URI)
	if err != nil {
		fmt.Printf("unable to create libvirt provider: %v", err)
		os.Exit(1)
	}
	file, err := i.GenerateFromInventory(p)
	if err != nil {
		fmt.Printf("unable to create libvirt provider: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(file))
}
