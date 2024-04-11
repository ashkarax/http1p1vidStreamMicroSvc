package handler

import (
	"api_gateway_svc/stream/pb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SendVideoStream(ctx *gin.Context, StreamClient pb.StreamServiceClient) {

	videochunks, err := StreamClient.SendVideoStream(context.Background(), &pb.VStreamRequest{})
	if err != nil {
		log.Printf("Error receiving video chunk: %v", err)
		ctx.HTML(500, "<h1>Internal Server Error</h1>", nil)
		return
	}
	defer videochunks.CloseSend()

	ctx.Header("Content-Type", "video/mp4")
	now := time.Now()
	fmt.Println("gateway recieving start time:", now.Format("15:04:05.000"))
	for {
		chunk, err := videochunks.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error receiving video chunk: %v", err)
			ctx.HTML(500, "<h1>Internal Server Error</h1>", nil)
			return
		}

		n, err := ctx.Writer.Write(chunk.Data)
		if err != nil {
			log.Printf("Error receiving video chunk: %v", err)
			fmt.Println("-------", n)
			ctx.HTML(500, "<h1>Internal Server Error</h1>", nil)
			return
		}

		//fmt.Println("streaming:", n)
		ctx.Writer.Flush()
	}
	fmt.Println("ApiGateway stream recieving end time:", now.Format("15:04:05.000"))

}
