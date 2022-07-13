package factory

import (
	"fmt"
	"terralist/pkg/auth"
	"terralist/pkg/auth/github"
)

func NewProvider(backend auth.Backend, config auth.Configurator) (auth.Provider, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("could not create a new provider with invalid configuration: %v", err)
	}

	config.SetDefaults()

	switch backend {
	case auth.GITHUB:
		creator := github.Creator{}
		return creator.New(config)
	default:
		return nil, fmt.Errorf("unrecognized backend type")
	}
}
