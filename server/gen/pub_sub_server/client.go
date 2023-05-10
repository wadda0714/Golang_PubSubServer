// Code generated by goa v3.11.3, DO NOT EDIT.
//
// PubSubServer client
//
// Command:
// $ goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server

package pubsubserver

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "PubSubServer" service client.
type Client struct {
	PublishEndpoint     goa.Endpoint
	SubscribeEndpoint   goa.Endpoint
	SendMessageEndpoint goa.Endpoint
}

// NewClient initializes a "PubSubServer" service client given the endpoints.
func NewClient(publish, subscribe, sendMessage goa.Endpoint) *Client {
	return &Client{
		PublishEndpoint:     publish,
		SubscribeEndpoint:   subscribe,
		SendMessageEndpoint: sendMessage,
	}
}

// Publish calls the "publish" endpoint of the "PubSubServer" service.
func (c *Client) Publish(ctx context.Context, p *PublishPayload) (res string, err error) {
	var ires any
	ires, err = c.PublishEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Subscribe calls the "subscribe" endpoint of the "PubSubServer" service.
func (c *Client) Subscribe(ctx context.Context, p *SubscribePayload) (res string, err error) {
	var ires any
	ires, err = c.SubscribeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// SendMessage calls the "sendMessage" endpoint of the "PubSubServer" service.
func (c *Client) SendMessage(ctx context.Context, p *SendMessagePayload) (res string, err error) {
	var ires any
	ires, err = c.SendMessageEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
