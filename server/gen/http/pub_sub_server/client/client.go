// Code generated by goa v3.2.5, DO NOT EDIT.
//
// PubSubServer client HTTP transport
//
// Command:
// $ goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the PubSubServer service endpoint HTTP clients.
type Client struct {
	// Publish Doer is the HTTP client used to make requests to the publish
	// endpoint.
	PublishDoer goahttp.Doer

	// Subscribe Doer is the HTTP client used to make requests to the subscribe
	// endpoint.
	SubscribeDoer goahttp.Doer

	// SendMessage Doer is the HTTP client used to make requests to the sendMessage
	// endpoint.
	SendMessageDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the PubSubServer service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PublishDoer:         doer,
		SubscribeDoer:       doer,
		SendMessageDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Publish returns an endpoint that makes HTTP requests to the PubSubServer
// service publish server.
func (c *Client) Publish() goa.Endpoint {
	var (
		decodeResponse = DecodePublishResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPublishRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PublishDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PubSubServer", "publish", err)
		}
		return decodeResponse(resp)
	}
}

// Subscribe returns an endpoint that makes HTTP requests to the PubSubServer
// service subscribe server.
func (c *Client) Subscribe() goa.Endpoint {
	var (
		decodeResponse = DecodeSubscribeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSubscribeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SubscribeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PubSubServer", "subscribe", err)
		}
		return decodeResponse(resp)
	}
}

// SendMessage returns an endpoint that makes HTTP requests to the PubSubServer
// service sendMessage server.
func (c *Client) SendMessage() goa.Endpoint {
	var (
		decodeResponse = DecodeSendMessageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSendMessageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SendMessageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PubSubServer", "sendMessage", err)
		}
		return decodeResponse(resp)
	}
}
