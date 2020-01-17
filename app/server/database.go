package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
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
	query := `CREATE TABLE IF NOT EXISTS firms
			(
				id INT PRIMARY KEY AUTO_INCREMENT,
				title VARCHAR(32) NOT NULL UNIQUE
			);
			CREATE TABLE IF NOT EXISTS marks
			(
				id INT PRIMARY KEY AUTO_INCREMENT,
				title VARCHAR(32) NOT NULL UNIQUE,
				firmId INT(11) NOT NULL,
				FOREIGN KEY (firmId)  REFERENCES firms (Id)
			);
			CREATE TABLE IF NOT EXISTS cars
			(
				id INT PRIMARY KEY AUTO_INCREMENT,
				firmId INT(11) NOT NULL,
				markId INT(11) NOT NULL,
				litresVolume FLOAT(11) NOT NULL,
				FOREIGN KEY (firmId)  REFERENCES firms (Id),
				FOREIGN KEY (markId)  REFERENCES marks (Id)
			);
			INSERT IGNORE INTO firms (title) VALUES ('BMW'), ('Audi'), ('Toyota');
			INSERT IGNORE INTO marks (title, firmId) VALUES ('X5', 1);
			INSERT IGNORE INTO cars (markId, firmId, litresVolume) VALUES (1, 1, 6.6);`
	requests := strings.Split(query, ";")

	for _, v := range requests {
		strings.TrimSpace(v)
		if v != "" {
			_, err = Db.Exec(v)
		}
	}

	return
}
