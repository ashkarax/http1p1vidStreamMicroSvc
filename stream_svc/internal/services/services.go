package services

import (
	"fmt"
	"io"
	"os"
	"stream_svc/internal/pb"
	"time"
)

type StreamServer struct {
	pb.StreamServiceServer
}

func (s *StreamServer) SendVideoStream(req *pb.VStreamRequest, stream pb.StreamService_SendVideoStreamServer) error {

	filePath := "/home/ashkar/Documents/vidStreamHTTP1P1/stream_svc/sampleVideo.mp4"
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("-------------", err)
		return fmt.Errorf("error checking file: %w", err)
	}

	videoStream, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening video stream: %w", err)
	}
	defer videoStream.Close()

	buffer := make([]byte, 1024)

	now := time.Now()
	fmt.Println("Stream service start time:", now.Format("15:04:05.000"))

	for {
		n, err := videoStream.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading video stream: %w", err)
		}

		if err := stream.Send(&pb.VStreamResponse{Data: buffer[:n]}); err != nil {
			return fmt.Errorf("error sending video chunk: %w", err)
		}
	}

	fmt.Println("Stream service end time:", now.Format("15:04:05.000"))

	return nil

}
