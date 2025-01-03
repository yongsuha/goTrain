package mysql

import (
	"context"
	"errors"
	"github.com/yongsuha/goTrain/train-rpc/pkg/db"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"gorm.io/gorm"
)

type TrainsConfModel struct{}

func NewTrainsConfModel() biz.TrainsConfRepo {
	return &TrainsConfModel{}
}

var (
	trainDB *gorm.DB
)

func (t *TrainsConfModel) CreateTrain() error {
	trainDB = db.GetDB()
	//trainDB.Create(biz.TrainsConf{})
	if trainDB == nil {
		return errors.New("the trains table init failed")
	}
	return nil
}

func (t *TrainsConfModel) AddTrain(ctx context.Context, train *biz.TrainsConf) (int64, error) {
	result := trainDB.WithContext(ctx).Create(train)
	if result.Error != nil {
		return 0, errors.New("add to trains failed")
	}
	return train.TrainId, nil
}

func (t *TrainsConfModel) DelTrain(ctx context.Context, trainId int64) error {
	result := trainDB.WithContext(ctx).Where("train_id =?", trainId).Delete(biz.TrainsConf{})
	if result.Error != nil {
		return errors.New("delete train record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no train record found for the given train_id")
	}
	return nil
}

func (t *TrainsConfModel) GetTrainDetail(ctx context.Context, trainId int64) (*biz.TrainsConf, error) {
	train := biz.TrainsConf{}
	result := trainDB.WithContext(ctx).Where("train_id =?", trainId).First(&train)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("train record not found")
		}
		return nil, errors.New("failed to get train detail")
	}
	return &train, nil
}

func (t *TrainsConfModel) UpdateTrain(ctx context.Context, train *biz.TrainsConf) error {
	result := trainDB.WithContext(ctx).Model(biz.TrainsConf{}).Where("train_id =?", train.TrainId).Updates(train)
	if result.Error != nil {
		return errors.New("update train record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no train record found for the given train_id")
	}
	return nil
}

func (t *TrainsConfModel) GetTrains(ctx context.Context, bigId int64, limit int) ([]*biz.TrainsConf, error) {
	trains := []*biz.TrainsConf{}
	result := trainDB.WithContext(ctx).Where("train_id >=?", bigId).Limit(limit).Find(&trains)
	if result.Error != nil {
		return nil, errors.New("failed to get trains")
	}
	return trains, nil
}
