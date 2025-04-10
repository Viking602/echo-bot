package data

import (
	"context"
	"echo/internal/biz"
	"echo/internal/data/ent/subbililive"
	"time"
)

type subBiliLiveRepo struct {
	data *Data
}

func NewSubBiliLiveRepo(data *Data) biz.SubBiliLiveRepo {
	return &subBiliLiveRepo{
		data: data,
	}
}

func (r *subBiliLiveRepo) Create(ctx context.Context, sub *biz.SubBiliLive) (int64, error) {
	create, err := r.data.db.SubBiliLive.Create().
		SetRoomID(sub.RoomId).
		SetLiveStartTime(sub.LiveStartTime).
		SetLiveEndTime(sub.LiveEndTime).
		SetLiveState(sub.LiveState).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return 0, err
	}

	return create.ID, nil
}

func (r *subBiliLiveRepo) GetByRoomId(ctx context.Context, roomId int64) (*biz.SubBiliLive, error) {
	first, err := r.data.db.SubBiliLive.Query().Where(subbililive.RoomID(roomId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.SubBiliLive{
		Id:            first.ID,
		RoomId:        first.RoomID,
		LiveStartTime: first.LiveStartTime,
		LiveEndTime:   first.LiveEndTime,
		LiveState:     first.LiveState,
	}, nil
}

func (r *subBiliLiveRepo) Get(ctx context.Context) ([]*biz.SubBiliLive, error) {
	var result []*biz.SubBiliLive
	subs, err := r.data.db.SubBiliLive.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, s := range subs {
		result = append(result, &biz.SubBiliLive{
			Id:            s.ID,
			RoomId:        s.RoomID,
			LiveStartTime: s.LiveStartTime,
			LiveEndTime:   s.LiveEndTime,
			LiveState:     s.LiveState,
		})
	}
	return result, nil
}

func (r *subBiliLiveRepo) UpdateByRoomId(ctx context.Context, live *biz.SubBiliLive) error {
	_, err := r.data.db.SubBiliLive.Update().
		Where(subbililive.RoomID(live.RoomId)).
		SetUpdateTime(time.Now()).
		SetLiveState(live.LiveState).
		SetLiveStartTime(live.LiveStartTime).
		SetLiveEndTime(live.LiveEndTime).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
