package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yongsuha/goTrain/train-rpc/pkg/redisCli"
	"time"
)

const (
	SetSeat  = "ss_%d_%s" // 座位余票信息缓存 车次id和座位类型
	SeatLeft = "sl_%d"    // 座位余票
)

type SeatsCache struct {
	redisObj *redis.Client
}

func NewSeatsCache() *SeatsCache {
	return &SeatsCache{redisObj: redisCli.GetRedis(0)}
}

func (s *SeatsCache) SetSeat(ctx context.Context, trainId int64, seatType string, seatLeft int64) error {
	key := fmt.Sprintf(SetSeat, trainId, seatType)
	value := fmt.Sprintf(SeatLeft, seatLeft)
	err := s.redisObj.Set(ctx, key, value, 300*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *SeatsCache) GetSeat(ctx context.Context, trainId int64, seatType string) (int64, error) {
	key := fmt.Sprintf(SetSeat, trainId, seatType)
	seats, err := s.redisObj.Get(ctx, key).Int64()
	if err != nil && err != redis.Nil {
		return -1, err
	}
	return seats, nil
}

func (s *SeatsCache) DelSeat(ctx context.Context, trainId int64, seatType string) error {
	key := fmt.Sprintf(SetSeat, trainId, seatType)
	err := s.redisObj.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *SeatsCache) DecrementSeat(ctx context.Context, trainId int64, seatType string, decreNum int64) error {
	key := fmt.Sprintf(SetSeat, trainId, seatType)
	err := s.redisObj.DecrBy(ctx, key, decreNum).Err()
	if err != nil {
		return err
	}
	return nil
}
