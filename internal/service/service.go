package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewSetMasterService,
	NewBiliLiveAddService,
	NewBiliLiveDelService,
	NewDouyuLiveAddService,
	NewDouyuLiveDelService,
	NewGroupMoyuService,
)
