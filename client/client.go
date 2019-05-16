package client

import (
	"context"

	pb "github.com/wwgberlin/grpc/salute"
	"google.golang.org/grpc"
)

//Send takes an address to dial to, a request and a context.
//it opens an (insecure) connection to the server using the address and the context
//it creates a saluter client with the connection created
//it salutes (calls Salute on the client) giving the context and request args.
//Follow example: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func Send(address string, req *pb.SaluteRequest, ctx context.Context) (*pb.SaluteResponse, error) {
	//implement here!

	//Don't forget to close the connection on exit!
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewSaluterClient(conn)

	return c.Salute(ctx, req)

}
