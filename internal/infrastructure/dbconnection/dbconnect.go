package dbconnection

import (
	"github.com/iancoleman/strcase"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func CreateDBConnection(dataSourceName string) (*sqlx.DB, error) {
	var err error
	DB, err := sqlx.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	//mapping between fields in struct and database (default lower case)_
	DB.MapperFunc(strcase.ToSnake)

	if err = DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil
}
