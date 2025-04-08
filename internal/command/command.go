package command

import (
	"echo/internal/model"
	"echo/internal/registry"
	"echo/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEchoCommand,
	NewInitializedRegistry,
)

func NewInitializedRegistry(echoService *service.EchoService) *registry.CommandRegistry {
	reg := registry.NewCommandRegistry()
	registrars := []model.CommandRegistrar{
		NewEchoCommand(echoService),
	}

	for _, registrar := range registrars {
		registrar.Register(reg)
	}
	return reg
}
