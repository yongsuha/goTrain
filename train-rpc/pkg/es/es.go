package es

import (
	"github.com/BurntSushi/toml"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

// 配置结构体，映射 config.toml 中的 Elasticsearch 配置
type Config struct {
	Elasticsearch struct {
		Address              string `toml:"address"`
		Timeout              int    `toml:"timeout"`
		MaxIdleConnections   int    `toml:"max_idle_connections"`
		MaxActiveConnections int    `toml:"max_active_connections"`
		ConnectionTimeout    int    `toml:"connection_request_timeout"`
	} `toml:"elasticsearch"`
}

var (
	esCli *elastic.Client
)

func GetESCli() *elastic.Client {
	if esCli != nil {
		return esCli
	}
	// 读取配置文件
	var config Config
	_, err := toml.DecodeFile("/Users/tal/go/goTrain/config/config.toml", &config)
	if err != nil {
		log.Fatalf("Failed to decode config.toml: %v", err)
	}
	// 初始化 Elasticsearch 客户端
	esCli, err = elastic.NewClient(
		elastic.SetURL(config.Elasticsearch.Address),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elastic.SetGzip(true),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(true),
	)
	if err != nil {
		log.Fatalf("Failed to connect to Elasticsearch: %v", err)
	}
	log.Println("Connected to Elasticsearch successfully.")
	return esCli
}
