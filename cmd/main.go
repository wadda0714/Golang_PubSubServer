package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		serverType = flag.String("role", "nothing", "server type")
		listenPort = flag.String("port", "5555", "listen port")
	)
	flag.Parse()

	fmt.Println("server type:", *serverType)
	fmt.Println("listen port:", *listenPort)

	switch *serverType {
	case "server":

	case "client":
	}

}
