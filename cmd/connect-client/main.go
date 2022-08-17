package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"

	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"github.com/bufbuild/connect-go"
)

var (
	addr = flag.String("addr", "localhost:50007", "the address to connect to")
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		&http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
					// If you're also using this client for non-h2c traffic, you may want to
					// delegate to tls.Dial if the network isn't TCP or the addr isn't in an
					// allowlist.
					return net.Dial(network, addr)
				},
			},
		},
		fmt.Sprintf("http://%s", *addr),
		connect.WithGRPC(),
	)

	req := connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"})
	req.Header().Set("Dapr-App-Id", "server")

	res, err := client.Greet(
		context.Background(),
		req,
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
