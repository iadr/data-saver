package container

import (
	"ValuesImporter/facade/environment"
	"ValuesImporter/persistence"
	"ValuesImporter/persistence/postgres"
	"time"
)

var ENVIRONMENT *environment.Environment

func CreateApp(ENV *environment.Environment) {
	ENVIRONMENT = ENV
	startPersistence()
	createUsecases()
	startWorkers()

	for {
		time.Sleep(5 * time.Second)
	}
}

func startPersistence() {
	poolConnection := postgres.NewPostgreSQL(ENVIRONMENT)
	createRepositories(poolConnection)

}

func startWorkers() {
}

func createRepositories(poolConnection *persistence.PoolConnection) {
}

func createUsecases() {
}
