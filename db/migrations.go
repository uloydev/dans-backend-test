package db

import (
	"dans-backend-test/db/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{

	migrations.CreateUsersTable,
}
