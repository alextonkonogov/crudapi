package car

import (
	"strconv"
	"strings"

	"github.com/alextonkonogov/crudapi/app/server"
)

type car struct {
	Id           int     `db:"id"`
	FirmId       int     `db:"firmId"`
	Mark         string  `db:"mark"`
	LitresVolume float64 `db:"litresVolume"`

	Firm string `db:"title"`
}

func GetCarById(carId string) (c car, err error) {
	query := `SELECT * FROM cars WHERE id = ?`
	err = server.Db.QueryRowx(query, carId).StructScan(&c)
	return
}

func GetAllCars() (cars []car, err error) {
	query := `SELECT * FROM cars AS c 
			LEFT JOIN 
			(
				SELECT id AS firmId, title FROM firms
			) AS f 
			ON c.firmId = f.firmId`
	err = server.Db.Select(&cars, query)
	return
}

func (c *car) Add() (err error) {
	query := `INSERT INTO cars (firmId, mark, litresVolume) VALUES (?, ?, ?)`
	_, err = server.Db.Exec(query, c.FirmId, c.Mark, c.LitresVolume)
	return
}

func (c *car) Update() (err error) {
	query := `UPDATE cars SET firmId = ?, mark = ?, litresVolume = ? WHERE id = ?`
	_, err = server.Db.Exec(query, c.FirmId, c.Mark, c.LitresVolume, c.Id)
	return
}

func (c *car) Delete() (err error) {
	query := `DELETE FROM cars WHERE id = ?`
	_, err = server.Db.Exec(query, c.Id)
	return
}

func NewCar(firmIdStr, mark, litresVolumeStr string) (c *car, err error) {
	firmId, err := strconv.Atoi(firmIdStr)
	if err != nil {
		return
	}

	litresVolume, err := strconv.ParseFloat(strings.TrimSpace(litresVolumeStr), 64)
	if err != nil {
		return
	}

	c = &car{FirmId: firmId, Mark: mark, LitresVolume: litresVolume}
	return
}
