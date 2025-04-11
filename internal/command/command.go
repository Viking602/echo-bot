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

func NewInitializedRegistry(
	setService *service.SetMasterService,
	biliLiveAddService *service.BiliLiveAddService,
	biliLiveDelService *service.BiliLiveDelService,
) *registry.CommandRegistry {
	reg := registry.NewCommandRegistry()
	registrars := []model.CommandRegistrar{
		NewSetCommand(setService),
		NewLiveCommand(biliLiveAddService, biliLiveDelService),
	}

	for _, registrar := range registrars {
		registrar.Register(reg)
	}
	return reg
}
