package postgre

import (
	"github.com/bonaventurabs/be-tokopedia-auction/config"
	"github.com/bonaventurabs/be-tokopedia-auction/pkg/postgre"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func NewDB(dbc config.DBConfig) *DB {
	db := postgre.ConnectDB(dbc)
	return &DB{
		db,
	}
}
