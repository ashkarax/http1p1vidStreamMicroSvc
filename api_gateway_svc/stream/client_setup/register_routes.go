package clientsetup

import (
	"api_gateway_svc/config"
	"api_gateway_svc/stream/handler"
	"api_gateway_svc/stream/pb"

	"github.com/gin-gonic/gin"
)

type StreamServiceClient struct {
	StreamClient pb.StreamServiceClient
}

func RegisterRoutes(r *gin.Engine, config *config.Config) {
	svc := &StreamServiceClient{
		StreamClient: InitStreamServiceClient(config),
	}

	routes := r.Group("/stream")
	{
		routes.GET("/video", svc.SendVideoStream)
	}

}

func (u *StreamServiceClient) SendVideoStream(ctx *gin.Context) {
	handler.SendVideoStream(ctx, u.StreamClient)

}
