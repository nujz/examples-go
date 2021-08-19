package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	pb "github.com/nujz/examples-go/grpc/echo/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	timer := time.NewTicker(time.Second)
	defer timer.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-timer.C:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			if res, err := client.Test(ctx, &pb.EchoRequest{Msg: "hello"}); err != nil {
				panic(err)
			} else {
				fmt.Println(res.GetMsg())
			}
		case <-interrupt:
			return
		}
	}
}
