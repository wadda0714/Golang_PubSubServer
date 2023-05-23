package pubsubserverrestapi

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/wadda0714/Golang_PubSubServer/pkg/userdb"
	pubsubserver "github.com/wadda0714/Golang_PubSubServer/server/gen/pub_sub_server"
	"github.com/wadda0714/Golang_PubSubServer/util"
	"log"
)

// PubSubServer service example implementation.
// The example methods log the requests and return zero values.
type pubSubServersrvc struct {
	logger *log.Logger
}

// NewPubSubServer returns the PubSubServer service implementation.
func NewPubSubServer(logger *log.Logger) pubsubserver.Service {
	return &pubSubServersrvc{logger}
}

// Publish implements publish.
func (s *pubSubServersrvc) Publish(ctx context.Context, p *pubsubserver.PublishPayload) (res string, err error) {
	s.logger.Print("pubSubServer.publish")
	options := []qm.QueryMod{}
	ctx = context.Background()
	db, err := sql.Open("postgres", "user=root dbname=userdb password=root sslmode=disable port=5430 host=localhost ")
	defer db.Close()
	if err != nil {
		s.logger.Print("db connection error")
	}

	queries := append(options, []qm.QueryMod{
		userdb.UserWhere.ID.EQ(1),
	}...)

	var user *userdb.User
	if user, err = userdb.Users(queries...).One(ctx, db); err != nil {
		s.logger.Print(err)
		if err == sql.ErrNoRows {

			s.logger.Print("no user")
			return
		}
	}
	s.logger.Print(user.Username)

	return user.Username, nil

}

// Subscribe implements subscribe.
func (s *pubSubServersrvc) Subscribe(ctx context.Context, p *pubsubserver.SubscribePayload) (res string, err error) {
	s.logger.Print("pubSubServer.subscribe")
	nameList, err := util.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	s.logger.Print(nameList)

	return
}

// SendMessage implements sendMessage.
func (s *pubSubServersrvc) SendMessage(ctx context.Context, p *pubsubserver.SendMessagePayload) (res string, err error) {
	s.logger.Print("pubSubServer.sendMessage")

	f, err := util.CreateFile("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\n" + *p.Message)

	return "ok", nil

}
