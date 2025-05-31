package sqlutil

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabaseSQL(
	host string,
	port string,
	user string,
	password string,
	database string,
	driver string,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host,
		user,
		password,
		database,
		port,
	)

	switch driver {
	case DRIVER_POSTGRES:
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			return nil, err
		}

		return db, nil
	case DRIVER_MYSQL:
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			return nil, err
		}

		return db, nil

	default:
		return nil, errors.New("invalid driver: " + driver)
	}
}

func DisconnectDatabaseSQL(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
