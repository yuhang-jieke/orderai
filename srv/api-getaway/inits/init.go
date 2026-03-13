package inits

import (
	"log"

	"github.com/yuhang-jieke/orderai/srv/api-getaway/basic/config"
	__ "github.com/yuhang-jieke/orderai/srv/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	GrpcInit()
}
func GrpcInit() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.OrderClient = __.NewEcommerceServiceClient(conn)
}
