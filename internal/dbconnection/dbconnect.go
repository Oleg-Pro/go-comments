package dbconnection

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func CreateDBConnection(dataSourceName string) (*sqlx.DB, error) {
	var err error
	DB, err := sqlx.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil
}
