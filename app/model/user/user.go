package user

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/alextonkonogov/crudapi/app/server"
)

type user struct {
	Id       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
}

func GetUserById(userId string) (u user, err error) {
	query := `SELECT * FROM users WHERE id = ?`
	err = server.Db.QueryRowx(query, userId).StructScan(&u)
	return
}

func (u *user) Add() error {
	query := `INSERT INTO users (login, password, name, surname) VALUES (?,?,?,?)`
	_, err := server.Db.Exec(query, u.Login, u.Password, u.Name, u.Surname)
	return err
}

func (u *user) Update() error {
	query := `UPDATE users SET login = ?, name = ?, surname = ?  WHERE id = ?`
	_, err := server.Db.Exec(query, u.Login, u.Name, u.Surname, u.Id)
	return err
}

func (u *user) Delete() error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := server.Db.Exec(query, u.Id)
	return err
}

func Login(login, password string) (u user, err error) {
	query := `SELECT id, login, name, surname FROM users WHERE login = ? AND password = ?`
	err = server.Db.QueryRowx(query, login, password).StructScan(&u)
	return u, err
}

func GetAllUsers() (users []user, err error) {
	query := `SELECT id, login, name, surname FROM users`
	err = server.Db.Select(&users, query)
	return
}

func IsAdmin(userId, table string) (admin bool, err error) {
	query := `SELECT EXISTS (SELECT * FROM users_rights WHERE userId = ? AND tableName = ?)`
	err = server.Db.QueryRowx(query, userId, table).Scan(&admin)
	return admin, err
}

func NewUser(login, password, name, surname string) *user {
	pass := md5.Sum([]byte(password))
	hashedPass := hex.EncodeToString(pass[:])
	return &user{Login: login, Password: hashedPass, Name: name, Surname: surname}
}
