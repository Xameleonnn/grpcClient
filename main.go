package main

import (
	"context"
	"flag"
	"fmt"
	tester "github.com/Xameleonnn/grpctester"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var addr = flag.String("serveraddr", "wyf:5300", "where to bang to")
	flag.Parse()
	fmt.Printf("Addr - %s\n", *addr)
	client, err := newClient(*addr)
	if err != nil {
		log.Fatalf("Couldnt make client, error - %v", err)
	}

	fmt.Println("made a client successfully")
	grpcReq := tester.HandshakeReq{
		HelloOut: "from client",
	}
	i := 0
	for {
		i++
		start := time.Now()
		deadline := start.Add(3 * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		resp, errHandshake := client.Handshake(ctx, &grpcReq)
		cancel()
		if errHandshake != nil && errHandshake.Error() == "rpc error: code = DeadlineExceeded desc = context deadline exceeded" {
			fmt.Printf("deadline exceeded, iteration#%d\n", i)
			continue
		} else if errHandshake != nil {
			log.Fatalf("Couldnt make handshake, error - %v", errHandshake)
		}

		fmt.Printf("request sent, time elapsed - %d\n", time.Since(start))
		fmt.Printf("Got grpc response, message - %s\n", resp.HelloBack)

		time.Sleep(1 * time.Second)
	}
}

func newClient(addr string) (client tester.HandshakerClient, err error) {
	//resolver.SetDefaultScheme("dns")
	connection, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client = tester.NewHandshakerClient(connection)
	return
}
