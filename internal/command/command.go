package command

import (
	"echo/internal/model"
	"echo/internal/registry"
	"echo/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewSetCommand,
	NewInitializedRegistry,
)

func NewInitializedRegistry(setService *service.SetMasterService) *registry.CommandRegistry {
	reg := registry.NewCommandRegistry()
	registrars := []model.CommandRegistrar{
		NewSetCommand(setService),
	}

	for _, registrar := range registrars {
		registrar.Register(reg)
	}
	return reg
}
