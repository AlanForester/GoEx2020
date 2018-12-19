package srv

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

var DB *db
type db struct {
	Instance *sql.DB
	Error error
}

func (d *db) Open() error {
	d.Instance, d.Error = sql.Open("postgres", "dbname=gop user=gop")
	if d.Error != nil {
		return d.Error
	}

	boil.SetDB(d.Instance)

	return nil
}

func init() {
	DB = new(db)
	DB.Open()
}