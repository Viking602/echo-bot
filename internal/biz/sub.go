package biz

import (
	"context"
	"echo/internal/data/ent"
)

type Sub struct {
	Id         int64
	SubType    int64
	GroupId    int64
	SubId      int64
	BotId      int64
	Status     int
	CreateTime string
	UpdateTime string
}

type SubRepo interface {
	Create(ctx context.Context, sub *Sub) error
	GetSubByGroupIdSubType(ctx context.Context, groupId int64, subType int64) ([]*Sub, error)
	GetSubBySubType(ctx context.Context, subType int64) ([]*Sub, error)
	GetSubBySubId(ctx context.Context, subId int64) ([]*Sub, error)
	GetSubBySubIdGroupIdSubTypeBotId(ctx context.Context, subId int64, groupId int64, subType int64, botId int64) (*Sub, error)
	DelBySubIdBotIdGroupId(ctx context.Context, bizSub *Sub) error
	GetAllBySubIdSubType(ctx context.Context, subId int64, subType int64) ([]*Sub, error)
}

type SubUsecase struct {
	repo SubRepo
}

func NewSubUsecase(subRepo SubRepo) *SubUsecase {
	return &SubUsecase{repo: subRepo}
}

func (u *SubUsecase) CreateSub(ctx context.Context, sub *Sub) (error, bool) {
	_, err := u.repo.GetSubBySubIdGroupIdSubTypeBotId(ctx, sub.SubId, sub.GroupId, sub.SubType, sub.BotId)
	if err != nil {
		if ent.IsNotFound(err) {
			return u.repo.Create(ctx, sub), false
		}
		return err, false
	}

	return nil, true
}

func (u *SubUsecase) GetSubBySubIdSubType(ctx context.Context, subId int64, subType int64) ([]*Sub, error) {
	return u.repo.GetAllBySubIdSubType(ctx, subId, subType)
}

func (u *SubUsecase) GetSubByGroupIdSubType(ctx context.Context, groupId int64, subType int64) ([]*Sub, error) {
	return u.repo.GetSubByGroupIdSubType(ctx, groupId, subType)
}

func (u *SubUsecase) GetSubBySubType(ctx context.Context, subType int64) ([]*Sub, error) {
	return u.repo.GetSubBySubType(ctx, subType)
}

func (u *SubUsecase) GetSubBySubId(ctx context.Context, subId int64) ([]*Sub, error) {
	return u.repo.GetSubBySubId(ctx, subId)
}

func (u *SubUsecase) DelSub(ctx context.Context, sub *Sub) error {
	get, err := u.repo.GetSubBySubIdGroupIdSubTypeBotId(ctx, sub.SubId, sub.GroupId, sub.SubType, sub.BotId)
	if err != nil {
		return err
	}
	return u.repo.DelBySubIdBotIdGroupId(ctx, get)

}

func (u *SubUsecase) GetAllBySubIdSubType(ctx context.Context, subId int64, subType int64) ([]*Sub, error) {
	return u.repo.GetAllBySubIdSubType(ctx, subId, subType)
}
