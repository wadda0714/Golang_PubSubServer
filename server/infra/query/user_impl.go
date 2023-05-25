package query

import (
	"context"
	"database/sql"
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
		userdb.MessageWhere.Roomid.EQ("aaa1"),
	}...)
	var err error
	var msg []*userdb.Message
	if msg, err = userdb.Messages(queries...).All(ctx, u.db); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	var msgs []*model.Message
	for _, v := range msg {
		msgs = append(msgs, transfer.ConvertMessages(v))
	}

	return msgs, nil
}
