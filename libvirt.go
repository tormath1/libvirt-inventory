package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	libvirt "libvirt.org/libvirt-go"
)

var (
	DEFAULT_CONNECT_URI = "qemu:///system"
)

type Domain interface {
	// GetName returns the name of a domain
	GetName() string

	// GetIP returns the IP address of a domain
	GetIP() string
}

type Inventory interface {
	ListDomains() ([]Domain, error)
}

type provider struct {
	conn *libvirt.Connect
}

type domain struct {
	name string
	ip   string
}

func (d *domain) GetName() string {
	return d.name
}

func (d *domain) GetIP() string {
	return d.ip
}

// NewProvider returns a new libvirt provider
func NewProvider(URI string) (*provider, error) {
	c, err := libvirt.NewConnect(URI)
	return &provider{
		conn: c,
	}, err
}

// ListDomains returns the list of libvirt domains
// on the host
func (p *provider) ListDomains() ([]Domain, error) {
	// return only running domains
	domains, err := p.conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_RUNNING)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list all domains")
	}
	list := make([]Domain, len(domains))
	for i, d := range domains {
		if name, err := d.GetName(); err == nil {
			list[i] = &domain{
				name: name,
				ip:   "n/a",
			}
		}
	}
	return list, nil
}

func main() {
	p, err := NewProvider(DEFAULT_CONNECT_URI)
	if err != nil {
		fmt.Printf("unable to create libvirt provider: %v", err)
		os.Exit(1)
	}
	domains, _ := p.ListDomains()
	for _, d := range domains {
		fmt.Printf("%s\n", d.GetName())
	}
}
