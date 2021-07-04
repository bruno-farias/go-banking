package helper

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func GetDbClient() *sqlx.DB {
	dbDriver := GetEnvVariable("DB_DRIVER")
	dbUser := GetEnvVariable("DB_USER")
	dbPassword := GetEnvVariable("DB_PASSWORD")
	dbHost := GetEnvVariable("DB_HOST")
	dbPort := GetEnvVariable("DB_PORT")
	dbDatabase := GetEnvVariable("DB_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase)

	client, err := sqlx.Open(dbDriver, dataSourceName)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}