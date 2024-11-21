package postgres

import (
	"ValuesImporter/facade/environment"
	"ValuesImporter/persistence"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

var postgres *Postgres

var lock = &sync.Mutex{}

type Postgres struct {
	poolConnection persistence.PoolConnection
}

func NewPostgreSQL(ENVIRONMENT *environment.Environment) *persistence.PoolConnection {
	if postgres == nil {
		lock.Lock()
		defer lock.Unlock()
		if postgres == nil {
			poolConnection := connect(ENVIRONMENT)
			postgres = &Postgres{poolConnection: *poolConnection}
			return poolConnection
		}

	}
	return &postgres.poolConnection
}

func GetConnection() *persistence.PoolConnection {
	return &postgres.poolConnection
}

func connect(ENVIRONMENT *environment.Environment) *persistence.PoolConnection {

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		ENVIRONMENT.Database.Host,
		ENVIRONMENT.Database.Port,
		ENVIRONMENT.Database.Username,
		ENVIRONMENT.Database.Password,
		ENVIRONMENT.Database.Database,
		ENVIRONMENT.Database.Schema)
	db, err := sql.Open(ENVIRONMENT.Database.Driver, dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(ENVIRONMENT.Database.MaxDatabaseConnections)
	db.SetMaxIdleConns(ENVIRONMENT.Database.MaxIdleConnections)

	poolConnection := persistence.NewPoolConnection(db)
	return poolConnection
}
