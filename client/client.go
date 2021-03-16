package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "100bao.com/m/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

const (
	// address     = "localhost:50051"
	defaultName = "hello"
)

func etcdDial(c *clientv3.Client, service string) (*grpc.ClientConn, error) {
	etcdResolver, err := resolver.NewBuilder(c)
	if err != nil {
		return nil, err
	}
	return grpc.Dial("etcd://author/"+service,
		grpc.WithResolvers(etcdResolver),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "round_robin"}`)))
}

func main() {
	cli, e1 := clientv3.NewFromURL("http://localhost:2379")
	if e1 != nil {
		log.Fatalf("connect etcd failed.")
	}
	log.Printf("etcd connected.")

	for i := 0; i < 10; i++ {
		conn, err := etcdDial(cli, "hello-service")
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		log.Printf("grpc connected.")
		c := pb.NewGreeterClient(conn)
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Hello(ctx, &pb.StrRequest{Origin: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		time.Sleep(time.Duration(2) * time.Second)
	}
}
