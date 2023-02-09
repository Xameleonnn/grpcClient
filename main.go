package main

import (
	"context"
	"fmt"
	tester "github.com/Xameleonnn/grpctester"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	host = "172.17.0.1:5300"
)

func newClient(addr string) (client tester.HandshakerClient, err error) {
	//resolver.SetDefaultScheme("dns")
	connection, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client = tester.NewHandshakerClient(connection)
	return
}

func main() {
	ctx := context.Background()
	client, err := newClient(host)
	if err != nil {
		log.Fatalf("Couldnt make client, error - %v", err)
	}

	fmt.Println("made a client successfully")
	grpcReq := tester.HandshakeReq{
		HelloOut: "from client",
	}

	for {
		start := time.Now()
		//deadline := start.Add(3 * time.Second)
		//cancel := func() { log.Fatal("Deadline exceeded") }
		//ctx, cancel = context.WithDeadline(context.Background(), deadline)
		//defer cancel()
		resp, errHandshake := client.Handshake(ctx, &grpcReq)
		if errHandshake != nil {
			log.Fatalf("Couldnt make handshake, error - %v", err)
		}

		fmt.Printf("request sent, time elapsed - %d\n", time.Since(start))
		fmt.Printf("Got grpc response, message - %s\n", resp.HelloBack)

		time.Sleep(1 * time.Second)
	}
}
