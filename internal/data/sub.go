package data

import (
	"context"
	"echo/internal/biz"
	"echo/internal/data/ent"
	"echo/internal/data/ent/sub"
	"time"
)

type subRepo struct {
	data *Data
}

func NewSubRepo(data *Data) biz.SubRepo {
	return &subRepo{data: data}
}

func (r *subRepo) Create(ctx context.Context, sub *biz.Sub) error {
	_, err := r.data.db.Sub.Create().
		SetSubID(sub.SubId).
		SetSubType(sub.SubType).
		SetGroupID(sub.GroupId).
		SetStatus(1).
		SetBotID(sub.BotId).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *subRepo) GetSubByGroupIdSubType(ctx context.Context, groupId int64, subType int64) ([]*biz.Sub, error) {
	var result []*biz.Sub
	subs, err := r.data.db.Sub.Query().Where(sub.GroupID(groupId), sub.SubType(subType)).All(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, s := range subs {
		result = append(result, &biz.Sub{
			Id:         s.ID,
			SubId:      s.SubID,
			SubType:    s.SubType,
			GroupId:    s.GroupID,
			Status:     s.Status,
			BotId:      s.BotID,
			CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}

func (r *subRepo) GetSubBySubType(ctx context.Context, subType int64) ([]*biz.Sub, error) {
	var result []*biz.Sub
	subs, err := r.data.db.Sub.Query().Where(sub.SubType(subType)).All(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	for _, s := range subs {
		result = append(result, &biz.Sub{
			Id:         s.ID,
			SubId:      s.SubID,
			SubType:    s.SubType,
			GroupId:    s.GroupID,
			Status:     s.Status,
			BotId:      s.BotID,
			CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}

func (r *subRepo) GetSubBySubId(ctx context.Context, subId int64) ([]*biz.Sub, error) {
	var result []*biz.Sub
	subs, err := r.data.db.Sub.Query().Where(sub.SubID(subId)).All(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	for _, s := range subs {
		result = append(result, &biz.Sub{
			Id:         s.ID,
			SubId:      s.SubID,
			SubType:    s.SubType,
			GroupId:    s.GroupID,
			Status:     s.Status,
			BotId:      s.BotID,
			CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil

}

func (r *subRepo) GetSubBySubIdGroupIdSubTypeBotId(ctx context.Context, subId int64, groupId int64, subType int64, botId int64) (*biz.Sub, error) {
	first, err := r.data.db.Sub.Query().Where(sub.SubID(subId), sub.GroupID(groupId), sub.SubType(subType), sub.BotID(botId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Sub{
		Id:         first.ID,
		SubId:      first.SubID,
		SubType:    first.SubType,
		GroupId:    first.GroupID,
		Status:     first.Status,
		BotId:      first.BotID,
		CreateTime: first.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: first.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *subRepo) DelBySubIdBotIdGroupId(ctx context.Context, bizSub *biz.Sub) error {
	_, err := r.data.db.Sub.Delete().Where(sub.SubID(bizSub.SubId), sub.BotID(bizSub.BotId), sub.GroupID(bizSub.GroupId)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil

}

func (r *subRepo) GetAllBySubIdSubType(ctx context.Context, subId int64, subType int64) ([]*biz.Sub, error) {
	var result []*biz.Sub

	subs, err := r.data.db.Sub.Query().Where(sub.SubID(subId), sub.SubType(subType)).All(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	for _, s := range subs {
		result = append(result, &biz.Sub{
			Id:         s.ID,
			SubId:      s.SubID,
			SubType:    s.SubType,
			GroupId:    s.GroupID,
			Status:     s.Status,
			BotId:      s.BotID,
			CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: s.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}
