package test

import (
	"github.com/adimax2953/host"
)

type MockComponent struct {
}

func (c *MockComponent) Runner() host.Runner {
	return &MockComponentRunner{
		prefix: "MockComponent",
	}
}
