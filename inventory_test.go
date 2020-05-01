package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeProvider struct{}

type fakeDomain struct {
	name string
	ip   string
}

func (fd *fakeDomain) GetName() string {
	return fd.name
}

func (fd *fakeDomain) GetIP() string {
	return fd.ip
}

func (fp *fakeProvider) ListDomains() ([]Domain, error) {
	return []Domain{
		&fakeDomain{
			name: "domain-1",
			ip:   "ip1",
		},
		&fakeDomain{
			name: "domain-2",
			ip:   "ip2",
		},
	}, nil
}

func TestInventoryGeneration(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var i InventoryFile
		p := &fakeProvider{}
		inv, err := i.GenerateFromInventory(p)
		require.Nil(t, err)
		assert.Equal(t, inv, []byte(`{"active":{"hosts":["ip1","ip2"],"vars":{}},"_meta":{"hostvars":null}}`))
	})
}
