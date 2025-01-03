package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/yongsuha/goTrain/train-rpc/pkg/db"
	"github.com/yongsuha/goTrain/train-rpc/pkg/redisCli"
	"github.com/yongsuha/goTrain/train-rpc/source/data/mysql"
	"log"
)

type Config struct {
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	} `toml:"server"`
	Database struct {
		Driver         string `toml:"driver"`
		Username       string `toml:"username"`
		Password       string `toml:"password"`
		Host           string `toml:"host"`
		Port           int    `toml:"port"`
		Name           string `toml:"name"`
		MaxConnections int    `toml:"max_connections"`
	} `toml:"database"`
	Cache struct {
		Type                 string `toml:"type"`
		Host                 string `toml:"host"`
		Port                 int    `toml:"port"`
		Password             string `toml:"password"`
		MaxIdleConnections   int    `toml:"max_idle_connections"`
		MaxActiveConnections int    `toml:"max_active_connections"`
	} `toml:"cache"`
	Log struct {
		Level string `toml:"level"`
		Path  string `toml:"path"`
	} `toml:"log"`
}

func main() {
	var config Config
	_, err := toml.DecodeFile("/Users/tal/go/goTrain/config/config.toml", &config)
	if err != nil {
		fmt.Println("读取配置文件错误:", err)
		return
	}
	fmt.Printf("服务器配置: %+v\n", config.Server)
	fmt.Printf("数据库配置: %+v\n", config.Database)
	fmt.Printf("缓存配置: %+v\n", config.Cache)
	fmt.Printf("日志配置: %+v\n", config.Log)

	db.GetDB()
	redisCli.GetRedis(0)
	err = mysql.NewTrainsConfModel().CreateTrain()
	if err != nil {
		log.Fatalln("the trains table init failed")
	}
	err = mysql.NewUsersConfModel().CreateUser()
	if err != nil {
		log.Fatalln("the users table init failed")
	}
	err = mysql.NewSeatsConfModel().CreateSeat()
	if err != nil {
		log.Fatalln("the seats table init failed")
	}
	err = mysql.NewOrdersConfModel().CreateOrder()
	if err != nil {
		log.Fatalln("the orders table init failed")
	}
	err = mysql.NewTicketsConfModel().CreateTicket()
	if err != nil {
		log.Fatalln("the tickets table init failed")
	}
	err = mysql.NewOrderTicketConfModel().CreateOrderTicket()
	if err != nil {
		log.Fatalln("the order_tickets table init failed")
	}
}
