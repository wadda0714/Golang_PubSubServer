// Code generated by goa v3.11.3, DO NOT EDIT.
//
// PubSubServer HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wadda0714/Golang_PubSubServer/server/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	pubsubserverc "github.com/wadda0714/Golang_PubSubServer/gen/http/pub_sub_server/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `pub-sub-server publish
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` pub-sub-server publish` + "\n" +
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
) (goa.Endpoint, any, error) {
	var (
		pubSubServerFlags = flag.NewFlagSet("pub-sub-server", flag.ContinueOnError)

		pubSubServerPublishFlags = flag.NewFlagSet("publish", flag.ExitOnError)
	)
	pubSubServerFlags.Usage = pubSubServerUsage
	pubSubServerPublishFlags.Usage = pubSubServerPublishUsage

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
		data     any
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
				data = nil
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
    %[1]s [globalflags] pub-sub-server COMMAND [flags]

COMMAND:
    publish: Publish implements publish.

Additional help:
    %[1]s pub-sub-server COMMAND --help
`, os.Args[0])
}
func pubSubServerPublishUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] pub-sub-server publish

Publish implements publish.

Example:
    %[1]s pub-sub-server publish
`, os.Args[0])
}
