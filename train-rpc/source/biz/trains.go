package biz

import (
	"context"
	"time"
)

const (
	TrainESMapping = `
{
	"mappings": {
		"properties": {
			"trainId": {"type": "integer"},
			"trainNumber": {"type": "text"},
			"departureStation": {"type": "text"},
			"arrivalStation": {"type": "text"},
			"totalSeat": {"type": "integer"}
		}
	}
}`
	TrainsIndex = "train_trains"
)

type TrainsConf struct {
	TrainId          int64     `json:"train_id" gorm:"column:train_id"`
	TrainNumber      string    `json:"train_number" gorm:"column:train_number"`
	DepartureStation string    `json:"departure_station" gorm:"column:departure_station"`
	ArrivalStation   string    `json:"arrival_station" gorm:"column:arrival_station"`
	TotalSeat        int64     `json:"total_seat" gorm:"column:total_seat"`
	DepartureTime    string    `json:"departure_time" gorm:"column:departure_time"`
	ArrivalTime      string    `json:"arrival_time" gorm:"column:arrival_time"`
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime       time.Time `json:"update_time" gorm:"column:update_time"`
}

type TrainsConfRepo interface {
	CreateTrain() error
	AddTrain(ctx context.Context, train *TrainsConf) (int64, error)
	UpdateTrain(ctx context.Context, train *TrainsConf) error
	GetTrainDetail(ctx context.Context, id int64) (*TrainsConf, error)
	DelTrain(ctx context.Context, id int64) error
	GetTrains(ctx context.Context, bigId int64, limit int) ([]*TrainsConf, error)
}

type EsTrains struct {
	TrainId          int64  `json:"train_id"`
	TrainNumber      string `json:"train_number"`
	DepartureStation string `json:"departure_station"`
	ArrivalStation   string `json:"arrival_station"`
	TotalSeat        int64  `json:"total_seat"`
}
