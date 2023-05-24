package query

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/wadda0714/Golang_PubSubServer/pkg/userdb"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/model"
	"github.com/wadda0714/Golang_PubSubServer/server/domain/query"
	"github.com/wadda0714/Golang_PubSubServer/server/transfer"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) query.UserQuery {
	return &User{db: db}
}

func (u *User) GetUsers(roomID int) ([]*model.User, error) {
	return nil, nil
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	return nil, nil
}

func (u *User) GetMessagesByRoomID(roomID int) ([]*model.Message, error) {
	options := []qm.QueryMod{}
	ctx := context.Background()

	queries := append(options, []qm.QueryMod{
		userdb.MessageWhere.Messageid.EQ("message1"),
	}...)
	var err error
	var msg *userdb.Message
	if msg, err = userdb.Messages(queries...).One(ctx, u.db); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	fmt.Println(msg.Content)
	var msgs []*model.Message
	//配列のmsgsにmsgを追加する
	msgs = append(msgs, transfer.ConvertMessages(msg))

	return msgs, nil
}
