package database

import (
	"fmt"
	"learn-grpc/internal/config"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(conf config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.Host,
		conf.User,
		conf.Password,
		conf.DBName,
		conf.Port,
		conf.SSLMode,
		conf.TimeZone,
	)
	var db *gorm.DB
	var err error
	retries := 10

	for i := range retries {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		if i == retries-1 {
			return nil, errors.Wrap(err, "unable to open gorm db")
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get sqlDB")
	}

	for i := range retries {
		err := sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retries-1 {
			return nil, errors.Wrap(err, "unable to ping")
		}
	}
	return db, nil
}
