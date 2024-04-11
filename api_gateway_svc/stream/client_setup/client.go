package clientsetup

import (
	"api_gateway_svc/config"
	"api_gateway_svc/stream/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitStreamServiceClient(config *config.Config) pb.StreamServiceClient {

	cc, err := grpc.NewClient(config.StreamSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect:", err)
	}

	return pb.NewStreamServiceClient(cc)
}
