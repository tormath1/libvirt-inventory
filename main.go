package main

import (
	"fmt"
	"os"
)

var (
	DEFAULT_CONNECT_URI = "qemu:///system"
)

func main() {
	URI := DEFAULT_CONNECT_URI
	if defaultURI := os.Getenv("VIRSH_DEFAULT_CONNECT_URI"); len(defaultURI) > 0 {
		URI = defaultURI
	}
	var i InventoryFile
	p, err := NewProvider(URI)
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
