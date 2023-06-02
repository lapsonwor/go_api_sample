package datastore

import (
	"fmt"
	"sync"
	"time"

	"lapson_go_api_sample/datastore/sqlc"
	"lapson_go_api_sample/pkg/logger"

	"lapson_go_api_sample/config"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Datastore struct {
	*sql.DB
	Queries *sqlc.Queries
	Mutex   *sync.RWMutex
}

var datastoreLogger *logrus.Entry = logger.GetLogger("datastore")
var DbQueries *sqlc.Queries

func NewPostgreSQL(dbConfig config.DatabaseConfig) *Datastore {
	var postgresqlConnectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.SSLMode)
	db, err := sql.Open("postgres", postgresqlConnectionString)
	if err != nil {
		datastoreLogger.Errorln("DB Connection error")
		panic(err)
	}
	if err := db.Ping(); err != nil {
		datastoreLogger.Errorln("Ping error")
		panic(err)
	}

	db.SetMaxIdleConns(dbConfig.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.MaxOpenConn)
	db.SetConnMaxLifetime(time.Hour)

	datastoreLogger.Infoln("Postgre connection established")
	DbQueries = sqlc.New(db)

	return &Datastore{
		db,
		DbQueries,
		&sync.RWMutex{},
	}
}
