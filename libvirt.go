package main

import (
	"strings"

	"github.com/pkg/errors"

	"libvirt.org/go/libvirt"
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

func (p *provider) getIP(domain libvirt.Domain) (string, error) {
	ifaces, err := domain.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
	if err != nil {
		return "n/a", errors.Wrapf(err, "unable to list all interface addresses: %v", err)
	}
	for _, iface := range ifaces {
		// TODO: handle different patterns
		if strings.Contains(iface.Name, "vnet") {
			if len(iface.Addrs) >= 1 {
				return iface.Addrs[0].Addr, nil
			}
		}
	}
	return "n/a", nil
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
		var (
			name string
			ip   string
		)
		if name, err = d.GetName(); err != nil {
			// TODO: display the error
			continue
		}
		ip, err = p.getIP(d)
		if err != nil {
			// TODO: display the error
			continue
		}
		list[i] = &domain{
			name: name,
			ip:   ip,
		}
	}
	return list, nil
}
