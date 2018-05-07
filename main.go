package main

import (
	"context"
	"plugin"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/PeerXu/error-grpc-with-plugin/proto"
)

func main() {
	lib, err := plugin.Open("./plugin.so")
	if err != nil {
		logrus.WithError(err).Fatalf("failed to open plugin")
	}
	fn, err := lib.Lookup("Serve")
	if err != nil {
		logrus.WithError(err).Fatalf("failed to lookup function")
	}
	go func() {
		<-time.After(3 * time.Second)
		conn, err := grpc.Dial("localhost:13401", grpc.WithInsecure())
		if err != nil {
			logrus.WithError(err).Fatalf("failed to dial to grpc")
		}
		defer conn.Close()
		cli := pb.NewGreetServiceClient(conn)
		res, err := cli.Greet(context.Background(), &pb.GreetRequest{})
		if err != nil {
			logrus.WithError(err).Fatalf("failed to call Greet")
		}
		logrus.Infof(res.Text)
	}()
	fn.(func())()
}
