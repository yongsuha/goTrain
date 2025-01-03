package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yongsuha/goTrain/train-rpc/pkg/redisCli"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"time"
)

const (
	TrainKey = "tk_%d" // 车次对应的缓存key 车次id
)

type TrainsCache struct {
	redisObj *redis.Client
}

func NewTrainsCache() *TrainsCache {
	return &TrainsCache{redisObj: redisCli.GetRedis(0)}
}

func (t *TrainsCache) BatchSetTrain(ctx context.Context, trains []*biz.TrainsConf) error {
	for _, train := range trains {
		key := fmt.Sprintf(TrainKey, train.TrainId)
		member, err := json.Marshal(train)
		if err != nil {
			return err
		}
		err = t.redisObj.Set(ctx, key, member, 24*time.Hour).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TrainsCache) BatchDelTrain(ctx context.Context, trains []int64) error {
	for _, train := range trains {
		key := fmt.Sprintf(TrainKey, train)
		err := t.redisObj.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TrainsCache) SetTrain(ctx context.Context, train *biz.TrainsConf) error {
	key := fmt.Sprintf(TrainKey, train.TrainId)
	member, err := json.Marshal(train)
	if err != nil {
		return err
	}
	err = t.redisObj.Set(ctx, key, member, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (t *TrainsCache) DelTrain(ctx context.Context, train int64) error {
	key := fmt.Sprintf(TrainKey, train)
	err := t.redisObj.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (t *TrainsCache) GetTrain(ctx context.Context, id int64) (*biz.TrainsConf, error) {
	key := fmt.Sprintf(TrainKey, id)
	train, err := t.redisObj.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if train == "" {
		return nil, nil
	}
	ret := &biz.TrainsConf{}
	err = json.Unmarshal([]byte(train), ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TrainsCache) BatchGetTrain(ctx context.Context, ids []int64) ([]*biz.TrainsConf, error) {
	ret := []*biz.TrainsConf{}
	for _, id := range ids {
		key := fmt.Sprintf(TrainKey, id)
		train, err := t.redisObj.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			return nil, err
		}
		if train == "" {
			return nil, nil
		}
		data := &biz.TrainsConf{}
		err = json.Unmarshal([]byte(train), data)
		if err != nil {
			return nil, err
		}
		ret = append(ret, data)
	}
	return ret, nil
}
