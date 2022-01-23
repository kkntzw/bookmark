package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kkntzw/bookmark/internal/di"
	"github.com/kkntzw/bookmark/internal/presentation/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Start")
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	bs := di.InjectBookmarkServer()
	pb.RegisterBookmarkerServer(s, bs)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	s.Stop()
	log.Println("Stop")
}
