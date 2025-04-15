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
	NewGroupCommand,
)

func NewInitializedRegistry(
	setService *service.SetMasterService,
	biliLiveAddService *service.BiliLiveAddService,
	biliLiveDelService *service.BiliLiveDelService,
	douyuLiveAddService *service.DouyuLiveAddService,
	douyuLiveDelService *service.DouyuLiveDelService,
	groupMoyuService *service.GroupMoyuService,
) *registry.CommandRegistry {
	reg := registry.NewCommandRegistry()
	registrars := []model.CommandRegistrar{
		NewSetCommand(setService),
		NewLiveCommand(biliLiveAddService, biliLiveDelService, douyuLiveAddService, douyuLiveDelService),
		NewGroupCommand(groupMoyuService),
	}

	for _, registrar := range registrars {
		registrar.Register(reg)
	}
	return reg
}
