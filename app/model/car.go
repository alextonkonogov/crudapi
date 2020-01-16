package model

import (
	"github.com/alextonkonogov/crudapi/app/server"
	"strconv"
	"strings"
)

type car struct {
	Id           int     `db:"id"`
	FirmId       int     `db:"firmId"`
	MarkId       int     `db:"markId"`
	LitresVolume float64 `db:"litresVolume"`
}

func GetCarById(carId string) (c car, err error) {
	query := `SELECT * FROM cars WHERE id = ?`
	err = server.Db.QueryRowx(query, carId).StructScan(&c)
	return
}

func GetAllCars() (cars []car, err error) {
	query := `SELECT * FROM cars`
	err = server.Db.Select(&cars, query)
	return
}

func (c *car) Add() (err error) {
	query := `INSERT INTO cars (firmId, markId, litresVolume) VALUES (?, ?, ?)`
	_, err = server.Db.Exec(query, c.FirmId, c.MarkId, c.LitresVolume)
	return
}

func (c *car) Update() (err error) {
	query := `UPDATE cars SET firmId = ?, markId = ?, litresVolume = ? WHERE id = ?`
	_, err = server.Db.Exec(query, c.FirmId, c.MarkId, c.LitresVolume, c.Id)
	return
}

func (c *car) Delete() (err error) {
	query := `DELETE FROM cars WHERE id = ?`
	_, err = server.Db.Exec(query, c.Id)
	return
}

func NewCar(firmIdStr, markIdStr, litresVolumeStr string) (c *car, err error) {
	firmId, err := strconv.Atoi(firmIdStr)
	if err != nil {
		return
	}
	markId, err := strconv.Atoi(markIdStr)
	if err != nil {
		return
	}

	litresVolume, err := strconv.ParseFloat(strings.TrimSpace(litresVolumeStr), 64)
	if err != nil {
		return
	}

	c = &car{FirmId: firmId, MarkId: markId, LitresVolume: litresVolume}
	return
}
