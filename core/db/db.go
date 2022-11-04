package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Config struct {
	Host          string
	Port          uint64
	Username      string
	Password      string
	DBName        string
	SSLMode       string
	TimeZone		string
	GormLogEnable bool
}

func New(cfg Config) *gorm.DB {
	var gormLogger logger.Interface

	if !cfg.GormLogEnable {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default
	}

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, cfg.TimeZone)), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
