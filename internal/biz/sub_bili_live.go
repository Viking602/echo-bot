package biz

import (
	"context"
	"echo/internal/data/ent"
	"time"
)

type SubBiliLive struct {
	Id            int64
	RoomId        int64
	LiveState     int64
	LiveStartTime int64
	LiveEndTime   int64
	CreateTime    time.Time
	UpdateTime    time.Time
}

type SubBiliLiveRepo interface {
	GetByRoomId(ctx context.Context, roomId int64) (*SubBiliLive, error)
	Create(ctx context.Context, sub *SubBiliLive) (int64, error)
	Get(ctx context.Context) ([]*SubBiliLive, error)
	UpdateByRoomId(ctx context.Context, live *SubBiliLive) error
	Del(ctx context.Context, id int64) error
}

type SubBiliLiveUsecase struct {
	repo SubBiliLiveRepo
}

func NewSubBiliLiveUsecase(repo SubBiliLiveRepo) *SubBiliLiveUsecase {
	return &SubBiliLiveUsecase{repo: repo}
}

func (u *SubBiliLiveUsecase) UpdateBiliLive(ctx context.Context, live *SubBiliLive) error {
	return u.repo.UpdateByRoomId(ctx, live)
}

func (u *SubBiliLiveUsecase) GetByRoomId(ctx context.Context, roomId int64) (*SubBiliLive, error) {
	return u.repo.GetByRoomId(ctx, roomId)
}

func (u *SubBiliLiveUsecase) CreateNewBiliLive(ctx context.Context, sub *SubBiliLive) (int64, error) {
	get, err := u.GetByRoomId(ctx, sub.RoomId)
	if err != nil {
		if ent.IsNotFound(err) {
			return u.repo.Create(ctx, sub)
		}
		return 0, err
	}

	return get.Id, nil
}

func (u *SubBiliLiveUsecase) GetAll(ctx context.Context) ([]*SubBiliLive, error) {
	return u.repo.Get(ctx)
}

func (u *SubBiliLiveUsecase) Del(ctx context.Context, id int64) error {
	return u.repo.Del(ctx, id)
}
