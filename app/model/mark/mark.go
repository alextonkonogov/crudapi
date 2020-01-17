package mark

import (
	"github.com/alextonkonogov/crudapi/app/server"
	"strconv"
)

type mark struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	FirmId int    `db:"firmId"`
}

func GetMarkById(markId string) (m mark, err error) {
	query := `SELECT * FROM marks WHERE id = ?`
	err = server.Db.QueryRowx(query, markId).StructScan(&m)
	return
}

func GetAllMarks() (marks []mark, err error) {
	query := `SELECT * FROM marks`
	err = server.Db.Select(&marks, query)
	return
}

func (m *mark) Add() (err error) {
	query := `INSERT INTO marks (title, firmId) VALUES (?,?)`
	_, err = server.Db.Exec(query, m.Title, m.FirmId)
	return
}

func (m *mark) Update() (err error) {
	query := `UPDATE marks SET title = ?, firmId = ? WHERE id = ?`
	_, err = server.Db.Exec(query, m.Title, m.FirmId, m.Id)
	return
}

func (m *mark) Delete() (err error) {
	query := `DELETE FROM marks WHERE id = ?`
	_, err = server.Db.Exec(query, m.Id)
	return
}

func NewMark(title, firmIdStr string) (m *mark, err error) {
	firmId, err := strconv.Atoi(firmIdStr)
	if err != nil {
		return
	}
	m = &mark{Title: title, FirmId: firmId}
	return
}
