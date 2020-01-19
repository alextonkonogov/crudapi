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
			CREATE TABLE IF NOT EXISTS cars
			(
				id INT PRIMARY KEY AUTO_INCREMENT,
				firmId INT(11) NOT NULL,
				mark VARCHAR(32) NOT NULL,
				litresVolume FLOAT(11) NOT NULL,
				FOREIGN KEY (firmId)  REFERENCES firms (Id),
				UNIQUE KEY uniq_index (firmId, mark, litresVolume)
			);
			CREATE TABLE IF NOT EXISTS users
			(
				id INT PRIMARY KEY AUTO_INCREMENT,
				login VARCHAR(32) NOT NULL UNIQUE,
				password VARCHAR(32) NOT NULL,
				name VARCHAR(32) NOT NULL,
				surname VARCHAR(32) NOT NULL
			);
			CREATE TABLE IF NOT EXISTS users_rights
			(
				id INT PRIMARY KEY AUTO_INCREMENT,
				userId INT(11) NOT NULL,
				tableName VARCHAR(32) NOT NULL,
				FOREIGN KEY (userId)  REFERENCES users (Id),
				UNIQUE KEY uniq_index2 (userId, tableName)
			);
			INSERT IGNORE INTO firms (title) VALUES ('BMW'), ('Audi'), ('Toyota');
			INSERT IGNORE INTO cars (firmId, mark, litresVolume) VALUES (1, 'X5', 6.6);
			INSERT IGNORE INTO users (login, password, name, surname) VALUES ('admin', '202cb962ac59075b964b07152d234b70', 'John', 'Snow');
			INSERT IGNORE INTO users_rights (userId, tableName) VALUES (1, 'users');`
	requests := strings.Split(query, ";")

	for _, v := range requests {
		strings.TrimSpace(v)
		if v != "" {
			_, err = Db.Exec(v)
		}
	}

	return
}
