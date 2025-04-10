package biz

import "context"

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
}

type SubUsecase struct {
	repo SubRepo
}

func NewSubUsecase(subRepo SubRepo) *SubUsecase {
	return &SubUsecase{repo: subRepo}
}

func (u *SubUsecase) CreateSub(ctx context.Context, sub *Sub) error {
	return u.repo.Create(ctx, sub)
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
