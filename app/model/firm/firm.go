package firm

import (
	"github.com/alextonkonogov/crudapi/app/server"
)

type firm struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
}

func GetFirmById(firmId string) (f firm, err error) {
	query := `SELECT * FROM firms WHERE id = ?`
	err = server.Db.QueryRowx(query, firmId).StructScan(&f)
	return
}

func GetAllFirms() (firms []firm, err error) {
	query := `SELECT * FROM firms`
	err = server.Db.Select(&firms, query)
	return
}

func (f *firm) Add() (err error) {
	query := `INSERT INTO firms (title) VALUES (?)`
	_, err = server.Db.Exec(query, f.Title)
	return
}

func (f *firm) Update() (err error) {
	query := `UPDATE firms SET title = ? WHERE id = ?`
	_, err = server.Db.Exec(query, f.Title, f.Id)
	return
}

func (f *firm) Delete() (err error) {
	query := `DELETE FROM firms WHERE id = ?`
	_, err = server.Db.Exec(query, f.Id)
	return
}

func NewFirm(title string) (f *firm) {
	return &firm{Title: title}
}
