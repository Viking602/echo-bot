package biz

import (
	"context"
	"echo/internal/data/ent"
	"time"
)

type SubDouyuLive struct {
	Id            int64
	RoomId        int64
	LiveState     int64
	LiveStartTime int64
	LiveEndTime   int64
	CreateTime    time.Time
	UpdateTime    time.Time
}

type SubDouyuLiveRepo interface {
	Create(ctx context.Context, sub *SubDouyuLive) (int64, error)
	Get(ctx context.Context) ([]*SubDouyuLive, error)
	Del(ctx context.Context, id int64) error
	UpdateByRoomId(ctx context.Context, live *SubDouyuLive) error
	GetByRoomId(ctx context.Context, roomId int64) (*SubDouyuLive, error)
}

type SubDouyuLiveUsecase struct {
	repo SubDouyuLiveRepo
}

func NewSubDouyuLiveUsecase(repo SubDouyuLiveRepo) *SubDouyuLiveUsecase {
	return &SubDouyuLiveUsecase{repo: repo}
}

func (u *SubDouyuLiveUsecase) CreateNewLive(ctx context.Context, sub *SubDouyuLive) (int64, error) {
	get, err := u.GetByRoomId(ctx, sub.RoomId)
	if err != nil {
		if ent.IsNotFound(err) {
			return u.repo.Create(ctx, sub)
		}
		return 0, err
	}

	return get.Id, nil
}

func (u *SubDouyuLiveUsecase) GetAll(ctx context.Context) ([]*SubDouyuLive, error) {
	return u.repo.Get(ctx)
}

func (u *SubDouyuLiveUsecase) Del(ctx context.Context, id int64) error {
	return u.repo.Del(ctx, id)
}

func (u *SubDouyuLiveUsecase) UpdateLive(ctx context.Context, live *SubDouyuLive) error {
	return u.repo.UpdateByRoomId(ctx, live)
}

func (u *SubDouyuLiveUsecase) GetByRoomId(ctx context.Context, roomId int64) (*SubDouyuLive, error) {
	return u.repo.GetByRoomId(ctx, roomId)
}
