package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() (err error) {
	var dataSourceName = "JZPT9dc4od:rRdia75axD@tcp(remotemysql.com:3306)/JZPT9dc4od"
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
	}

	err = initTables()
	return err
}

func initTables() (err error) {
	query := `
	CREATE TABLE IF NOT EXISTS cars
	(
		id int auto_increment primary key unique,
		firmId int(11) not null,
		markId int(11) not null,
		litresVolume float(11) not null
	);
	`
	_, err = Db.Exec(query)

	return
}
