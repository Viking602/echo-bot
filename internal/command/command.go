package command

import (
	"echo/internal/model"
	"echo/internal/registry"
	"echo/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewSetCommand,
	NewLiveCommand,
	NewInitializedRegistry,
)

func NewInitializedRegistry(setService *service.SetMasterService, liveService *service.LiveAddService) *registry.CommandRegistry {
	reg := registry.NewCommandRegistry()
	registrars := []model.CommandRegistrar{
		NewSetCommand(setService),
		NewLiveCommand(liveService),
	}

	for _, registrar := range registrars {
		registrar.Register(reg)
	}
	return reg
}
