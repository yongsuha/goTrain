package es

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/olivere/elastic/v7"
	"github.com/yongsuha/goTrain/train-rpc/pkg/es"
)

type EsCli struct {
	esClient *elastic.Client
}

func NewEsCli() *EsCli {
	return &EsCli{
		esClient: es.GetESCli(),
	}
}

func (e *EsCli) CreateIndex(ctx context.Context, index, bodyStr string) error {
	exists, err := e.esClient.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	svc := e.esClient.CreateIndex(index)
	if bodyStr != "" {
		svc.BodyString(bodyStr)
	}
	_, err = svc.Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsCli) DeleteIndex(ctx context.Context, index string) error {
	exists, err := e.esClient.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	svc := e.esClient.DeleteIndex(index)
	_, err = svc.Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsCli) SearchDocument(ctx context.Context, index string, searchSource *elastic.SearchSource) (int64, []*elastic.SearchHit, error) {
	searchResult, err := e.esClient.Search().Index(index).SearchSource(searchSource).Do(ctx)
	if err != nil {
		return 0, nil, err
	}
	total := searchResult.Hits.TotalHits.Value
	hits := searchResult.Hits.Hits
	return total, hits, nil
}

func (e *EsCli) InsertDocument(ctx context.Context, index, id string, doc interface{}) error {
	_, err := e.esClient.Index().Index(index).Id(id).BodyJson(doc).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsCli) UpdateDocumentById(ctx context.Context, index, id string, doc interface{}) error {
	_, err := e.esClient.Update().Index(index).Id(id).Doc(doc).Do(ctx)
	spew.Dump(err)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsCli) GetDocumentById(ctx context.Context, index, id string) (*elastic.GetResult, error) {
	result, err := e.esClient.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *EsCli) DeleteDocumentById(ctx context.Context, index, id string) error {
	_, err := e.esClient.Delete().Index(index).Id(id).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsCli) InsertBatchDocument(ctx context.Context, index string, docMaps []map[string]interface{}) error {
	bulkRequest := e.esClient.Bulk()
	for _, docMap := range docMaps {
		id := docMap["id"].(string)
		doc := docMap["doc"]
		req := elastic.NewBulkIndexRequest().Index(index).Id(id).Doc(doc)
		bulkRequest = bulkRequest.Add(req)
	}
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EsCli) DeleteBatchDocument(ctx context.Context, index string, ids []string) error {
	bulkRequest := e.esClient.Bulk()
	for _, id := range ids {
		req := elastic.NewBulkDeleteRequest().Index(index).Id(id)
		bulkRequest = bulkRequest.Add(req)
	}
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
