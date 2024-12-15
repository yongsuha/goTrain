package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Database struct {
		Driver   string `toml:"driver"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Name     string `toml:"name"`
	} `toml:"database"`
}

var (
	db *gorm.DB
)

func initDB() error {
	var config Config
	_, err := toml.DecodeFile("/Users/tal/go/goTrain/config/config.toml", &config)
	if err != nil {
		return fmt.Errorf("读取配置文件错误: %v", err)
	}
	// 构建 MySQL 连接字符串，这里需要替换为实际的 IP 地址、端口号、数据库名称、用户名和密码
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	// 使用 mysql.Open 打开数据库连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("the database connect successfully")
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		if err := initDB(); err != nil {
			log.Fatalf("数据库初始化错误: %v", err)
		}
	}
	return db
}
