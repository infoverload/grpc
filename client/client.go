package client

import (
	"context"
	"log"

	pb "github.com/wwgberlin/grpc/salute"
	g "google.golang.org/grpc"
)

//Send takes an address to dial to, a request and a context.
//it opens an (insecure) connection to the server using the address and the context
//it creates a saluter client with the connection created
//it salutes (calls Salute on the client) giving the context and request args.
//Follow example: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func Send(address string, req *pb.SaluteRequest, ctx context.Context) (*pb.SaluteResponse, error) {
	// Don't forget to close the connection on exit!
	// Set up a connection to the server.
	conn, err := g.Dial(address, g.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSaluterClient(conn)

	r, err := c.Salute(ctx, req)
	return r, err

}
