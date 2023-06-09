// Code generated by goa v3.2.5, DO NOT EDIT.
//
// PubSubServer HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	pubsubserverc "github.com/wadda0714/Golang_PubSubServer/server/gen/http/pub_sub_server/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `pub-sub-server (publish|subscribe|send-message)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` pub-sub-server publish --room-name "room1"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		pubSubServerFlags = flag.NewFlagSet("pub-sub-server", flag.ContinueOnError)

		pubSubServerPublishFlags        = flag.NewFlagSet("publish", flag.ExitOnError)
		pubSubServerPublishRoomNameFlag = pubSubServerPublishFlags.String("room-name", "REQUIRED", "roomName to publish")

		pubSubServerSubscribeFlags        = flag.NewFlagSet("subscribe", flag.ExitOnError)
		pubSubServerSubscribeRoomNameFlag = pubSubServerSubscribeFlags.String("room-name", "REQUIRED", "roomName to subscribe")

		pubSubServerSendMessageFlags        = flag.NewFlagSet("send-message", flag.ExitOnError)
		pubSubServerSendMessageRoomNameFlag = pubSubServerSendMessageFlags.String("room-name", "REQUIRED", "roomName to publish")
		pubSubServerSendMessageMessageFlag  = pubSubServerSendMessageFlags.String("message", "REQUIRED", "message to publish")
	)
	pubSubServerFlags.Usage = pubSubServerUsage
	pubSubServerPublishFlags.Usage = pubSubServerPublishUsage
	pubSubServerSubscribeFlags.Usage = pubSubServerSubscribeUsage
	pubSubServerSendMessageFlags.Usage = pubSubServerSendMessageUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "pub-sub-server":
			svcf = pubSubServerFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "pub-sub-server":
			switch epn {
			case "publish":
				epf = pubSubServerPublishFlags

			case "subscribe":
				epf = pubSubServerSubscribeFlags

			case "send-message":
				epf = pubSubServerSendMessageFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "pub-sub-server":
			c := pubsubserverc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "publish":
				endpoint = c.Publish()
				data, err = pubsubserverc.BuildPublishPayload(*pubSubServerPublishRoomNameFlag)
			case "subscribe":
				endpoint = c.Subscribe()
				data, err = pubsubserverc.BuildSubscribePayload(*pubSubServerSubscribeRoomNameFlag)
			case "send-message":
				endpoint = c.SendMessage()
				data, err = pubsubserverc.BuildSendMessagePayload(*pubSubServerSendMessageRoomNameFlag, *pubSubServerSendMessageMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// pub-sub-serverUsage displays the usage of the pub-sub-server command and its
// subcommands.
func pubSubServerUsage() {
	fmt.Fprintf(os.Stderr, `The PubSubServer service performs operations on messages.
Usage:
    %s [globalflags] pub-sub-server COMMAND [flags]

COMMAND:
    publish: Publish implements publish.
    subscribe: Subscribe implements subscribe.
    send-message: SendMessage implements sendMessage.

Additional help:
    %s pub-sub-server COMMAND --help
`, os.Args[0], os.Args[0])
}
func pubSubServerPublishUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] pub-sub-server publish -room-name STRING

Publish implements publish.
    -room-name STRING: roomName to publish

Example:
    `+os.Args[0]+` pub-sub-server publish --room-name "room1"
`, os.Args[0])
}

func pubSubServerSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] pub-sub-server subscribe -room-name STRING

Subscribe implements subscribe.
    -room-name STRING: roomName to subscribe

Example:
    `+os.Args[0]+` pub-sub-server subscribe --room-name "room1"
`, os.Args[0])
}

func pubSubServerSendMessageUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] pub-sub-server send-message -room-name STRING -message STRING

SendMessage implements sendMessage.
    -room-name STRING: roomName to publish
    -message STRING: message to publish

Example:
    `+os.Args[0]+` pub-sub-server send-message --room-name "room1" --message "hello"
`, os.Args[0])
}
