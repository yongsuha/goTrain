package mysql

import (
	"context"
	"errors"
	"github.com/yongsuha/goTrain/train-rpc/pkg/db"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"gorm.io/gorm"
)

type SeatsConfModel struct{}

func NewSeatsConfModel() biz.SeatsConfRepo {
	return &SeatsConfModel{}
}

var (
	seatDB *gorm.DB
)

func (s *SeatsConfModel) CreateSeat() error {
	seatDB = db.GetDB()
	//seatDB.Create(biz.UsersConf{})
	if seatDB == nil {
		return errors.New("the seats table init failed")
	}
	return nil
}

func (s *SeatsConfModel) AddSeat(ctx context.Context, seat *biz.SeatsConf) (int64, error) {
	result := seatDB.WithContext(ctx).Create(seat)
	if result.Error != nil {
		return 0, errors.New("add to seats failed")
	}
	return seat.SeatId, nil
}

func (s *SeatsConfModel) DelSeat(ctx context.Context, seatId int64) error {
	result := seatDB.WithContext(ctx).Where("seat_id =?", seatId).Delete(biz.SeatsConf{})
	if result.Error != nil {
		return errors.New("delete seat record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no seat record found for the given seat_id")
	}
	return nil
}

func (s *SeatsConfModel) GetSeatDetail(ctx context.Context, getDetailParam *biz.GetSeatDetailParam) (*biz.SeatsConf, error) {
	seat := biz.SeatsConf{}
	query := seatDB.WithContext(ctx)

	if getDetailParam.TrainId != 0 {
		query = query.Where("train_id =?", getDetailParam.TrainId)
	}
	if getDetailParam.SeatNumber != "" {
		query = query.Where("seat_number =?", getDetailParam.SeatNumber)
	}

	result := query.First(&seat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("seat record not found")
		}
		return nil, errors.New("failed to get seat detail")
	}
	return &seat, nil
}

func (s *SeatsConfModel) UpdateSeat(ctx context.Context, seat *biz.SeatsConf) error {
	result := seatDB.WithContext(ctx).Model(biz.SeatsConf{}).Where("seat_id =?", seat.SeatId).Updates(seat)
	if result.Error != nil {
		return errors.New("update seat record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no seat record found for the given seat_id")
	}
	return nil
}
