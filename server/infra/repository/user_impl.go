package repository

import (
	"database/sql"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/repository"
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repository.UserRepository {
	return &user{db}
}

func (u *user) AddUserToDB(user *model.User) (err error) {
	_, err = u.db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) DeleteUserFromDB(user *model.User) (err error) {
	_, err = u.db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (u *user) AddMessageToDB(message *model.Message) (err error) {
	return nil
}
