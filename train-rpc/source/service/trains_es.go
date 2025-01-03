package service

import (
	"context"
	"github.com/spf13/cast"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"github.com/yongsuha/goTrain/train-rpc/source/data/es"
	"github.com/yongsuha/goTrain/train-rpc/source/data/mysql"
	"log"
	"time"
)

type TrainsEsService struct {
	esCli      *es.EsCli
	trainsRepo biz.TrainsConfRepo
}

func NewTrainsEsService() *TrainsEsService {
	return &TrainsEsService{
		esCli:      es.NewEsCli(),
		trainsRepo: mysql.NewTrainsConfModel(),
	}
}

func (t *TrainsEsService) DelIndex(ctx context.Context) error {
	return t.esCli.DeleteIndex(ctx, biz.TrainsIndex)
}

func (t *TrainsEsService) DelOneToEs(ctx context.Context, trainId string) error {
	return t.esCli.DeleteDocumentById(ctx, biz.TrainsIndex, trainId)
}

func (t *TrainsEsService) SyncOneToEs(ctx context.Context, trainId string) error {
	tag := "TrainsEsService:SyncOneToEs"
	// 从数据库获取车次信息
	train, err := t.trainsRepo.GetTrainDetail(ctx, cast.ToInt64(trainId))
	if err != nil {
		log.Println(ctx, tag, "GetTrainDetail error:%s", err.Error())
		return err
	}
	if train == nil {
		log.Println("the train information is empty")
		return err
	}
	// 更新es
	esTrain := biz.EsTrains{
		TrainId:          train.TrainId,
		TrainNumber:      train.TrainNumber,
		DepartureStation: train.DepartureStation,
		ArrivalStation:   train.ArrivalStation,
		TotalSeat:        train.TotalSeat,
	}
	err = t.esCli.DeleteDocumentById(ctx, biz.TrainsIndex, trainId)
	if err != nil {
		log.Println(ctx, tag, "DeleteDocumentById error:%s", err.Error())
		return err
	}
	err = t.esCli.InsertDocument(ctx, biz.TrainsIndex, cast.ToString(esTrain.TrainId), esTrain)
	if err != nil {
		log.Println(ctx, tag, "InsertDocument error:%s", err.Error())
		return err
	}
	return nil
}

func (t *TrainsEsService) SyncAllToEs(ctx context.Context) error {
	tag := "TrainsEsService:SyncAllToEs"
	// 创建索引
	err := t.esCli.CreateIndex(ctx, biz.TrainsIndex, biz.TrainESMapping)
	if err != nil {
		log.Println(ctx, tag, "CreateIndex error:%s", err.Error())
		return err
	}
	// 从数据库获取trains信息
	var curId int64 = 0
	limit := 200
	for {
		trains, err := t.trainsRepo.GetTrains(ctx, curId, limit)
		if err != nil {
			log.Println(ctx, tag, "GetTrains error:%s", err.Error())
			return err
		}
		if len(trains) == 0 {
			log.Println(ctx, tag, "sync trains to es finished, curId: %d", curId)
			break
		}
		curId = trains[len(trains)-1].TrainId
		// 数据转换
		esIds := make([]string, 0)
		docMaps := []map[string]interface{}{}
		for _, train := range trains {
			trainDoc := &biz.EsTrains{
				TrainId:          train.TrainId,
				TrainNumber:      train.TrainNumber,
				DepartureStation: train.DepartureStation,
				ArrivalStation:   train.ArrivalStation,
				TotalSeat:        train.TotalSeat,
			}
			docMap := map[string]interface{}{
				"id":  cast.ToString(train.TrainId),
				"doc": trainDoc,
			}
			docMaps = append(docMaps, docMap)
			esIds = append(esIds, cast.ToString(train.TrainId))
		}
		// 删除已有数据
		err = t.esCli.DeleteBatchDocument(ctx, biz.TrainsIndex, esIds)
		if err != nil {
			log.Println(ctx, tag, "DeleteBatchDocument error:%s", err.Error())
			return err
		}
		// 写入es
		err = t.esCli.InsertBatchDocument(ctx, biz.TrainsIndex, docMaps)
		if err != nil {
			log.Println(ctx, tag, "InsertBatchDocument error:%s", err.Error())
			return err
		}
	}
	return nil
}

func (t *TrainsEsService) BatchSyncToEs(ctx context.Context, bigId string) error {
	tag := "TrainsEsService:BatchSyncToEs"
	curId := cast.ToInt64(bigId)
	limit := 200
	for {
		trains, err := t.trainsRepo.GetTrains(ctx, curId, limit)
		if err != nil {
			log.Println(ctx, tag, "GetTrains error:%s", err.Error())
			return err
		}
		if len(trains) == 0 {
			log.Println(ctx, tag, "sync trains to es finished, curId: %d", curId)
			break
		}
		curId = trains[len(trains)-1].TrainId
		// 数据转换
		esIds := make([]string, 0)
		docMaps := []map[string]interface{}{}
		for _, train := range trains {
			trainDoc := &biz.EsTrains{
				TrainId:          train.TrainId,
				TrainNumber:      train.TrainNumber,
				DepartureStation: train.DepartureStation,
				ArrivalStation:   train.ArrivalStation,
				TotalSeat:        train.TotalSeat,
			}
			docMap := map[string]interface{}{
				"id":  cast.ToString(train.TrainId),
				"doc": trainDoc,
			}
			docMaps = append(docMaps, docMap)
			esIds = append(esIds, cast.ToString(train.TrainId))
		}
		// 删除已有数据
		err = t.esCli.DeleteBatchDocument(ctx, biz.TrainsIndex, esIds)
		if err != nil {
			log.Println(ctx, tag, "DeleteBatchDocument error:%s", err.Error())
			return err
		}
		// 写入es
		err = t.esCli.InsertBatchDocument(ctx, biz.TrainsIndex, docMaps)
		if err != nil {
			log.Println(ctx, tag, "InsertBatchDocument error:%s", err.Error())
			return err
		}
	}
	return nil
}

func (t *TrainsEsService) SyncChangeToEs(ctx context.Context, trainId int64) error {
	tag := "TrainsEsService:SyncChangeToEs"
	time.Sleep(time.Second * 10)
	// 获取车次信息
	train, err := t.trainsRepo.GetTrainDetail(ctx, trainId)
	if err != nil {
		log.Println(ctx, tag, "GetTrainDetail error:%s", err.Error())
		return err
	}
	if train == nil {
		log.Println(ctx, tag, "the train information is empty")
		return err
	}
	// 更新es
	esTrain := &biz.EsTrains{
		TrainId:          train.TrainId,
		TrainNumber:      train.TrainNumber,
		DepartureStation: train.DepartureStation,
		ArrivalStation:   train.ArrivalStation,
		TotalSeat:        train.TotalSeat,
	}
	err = t.esCli.UpdateDocumentById(ctx, biz.TrainsIndex, cast.ToString(trainId), esTrain)
	if err != nil {
		log.Println(ctx, tag, "UpdateDocumentById error:%s", err.Error())
		return err
	}
	return nil
}

//func (t *TrainsEsService) SearchTrain(ctx context.Context, req *trainsESRpc.SearchTrainsReq) (*trainsESRpc.SearchTrainsResp, error) {
//
//}
