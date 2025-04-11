package data

import (
	"context"
	"echo/internal/biz"
	"echo/internal/data/ent/subdouyulive"
	"time"
)

type subDouyuLiveRepo struct {
	data *Data
}

func NewSubDouyuLiveRepo(data *Data) biz.SubDouyuLiveRepo {
	return &subDouyuLiveRepo{
		data: data,
	}
}

func (r *subDouyuLiveRepo) Create(ctx context.Context, sub *biz.SubDouyuLive) (int64, error) {
	create, err := r.data.db.SubDouyuLive.Create().
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

func (r *subDouyuLiveRepo) GetByRoomId(ctx context.Context, roomId int64) (*biz.SubDouyuLive, error) {
	first, err := r.data.db.SubDouyuLive.Query().Where(subdouyulive.RoomID(roomId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.SubDouyuLive{
		Id:            first.ID,
		RoomId:        first.RoomID,
		LiveStartTime: first.LiveStartTime,
		LiveEndTime:   first.LiveEndTime,
		LiveState:     first.LiveState,
	}, nil
}

func (r *subDouyuLiveRepo) UpdateByRoomId(ctx context.Context, live *biz.SubDouyuLive) error {
	_, err := r.data.db.SubDouyuLive.Update().
		Where(subdouyulive.RoomID(live.RoomId)).
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

func (r *subDouyuLiveRepo) Del(ctx context.Context, id int64) error {
	_, err := r.data.db.SubDouyuLive.Delete().Where(subdouyulive.ID(id)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil

}

func (r *subDouyuLiveRepo) Get(ctx context.Context) ([]*biz.SubDouyuLive, error) {
	var result []*biz.SubDouyuLive
	subs, err := r.data.db.SubDouyuLive.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, s := range subs {
		result = append(result, &biz.SubDouyuLive{
			Id:            s.ID,
			RoomId:        s.RoomID,
			LiveStartTime: s.LiveStartTime,
			LiveEndTime:   s.LiveEndTime,
			LiveState:     s.LiveState,
		})
	}
	return result, nil
}
