package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/linggaaskaedo/go-playground-wire/model/common"
)

type DBOptions struct {
	Config common.Configuration
}

func NewDB(opts *DBOptions) (*sql.DB, error) {
	databaseURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		opts.Config.Database.DBUser,
		opts.Config.Database.DBPassword,
		opts.Config.Database.DBHost,
		opts.Config.Database.DBPort,
		opts.Config.Database.DBName,
	)

	db, err := sql.Open("mysql", databaseURI)
	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(opts.Config.Database.DBConnMaxIdleTime)
	db.SetMaxOpenConns(opts.Config.Database.DBMaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(opts.Config.Database.DBConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(opts.Config.Database.DBConnMaxIdleTime) * time.Minute)

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
