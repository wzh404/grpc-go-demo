package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "100bao.com/m/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	endpoints "go.etcd.io/etcd/client/v3/naming/endpoints"
	grpc "google.golang.org/grpc"
)

// const (
// 	port string = ":50052"
// )
var port string = ":50051"

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.GreeterServer
}

func (s *server) Hello(ctx context.Context, in *pb.StrRequest) (*pb.StrReply, error) {
	log.Println("Received: " + in.GetOrigin())
	res := in.GetOrigin() + "-" + port
	return &pb.StrReply{Message: res}, nil
}

func etcdAdd(c *clientv3.Client, service, addr string) error {
	em, _ := endpoints.NewManager(c, service)
	return em.AddEndpoint(c.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr: addr})
}

func etcdDelete(c *clientv3.Client, service, addr string) error {
	em, _ := endpoints.NewManager(c, service)
	return em.DeleteEndpoint(c.Ctx(), service+"/"+addr)
}

func exitFunc() {
	os.Exit(0)
}

func main() {
	flag.StringVar(&port, "port", "50051", "端口")
	flag.Parse()

	cli, e1 := clientv3.NewFromURL("http://localhost:2379")
	if e1 != nil {
		log.Fatalf("connect etcd failed.")
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for s := range sc {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("退出", s)
				etcdDelete(cli, "hello-service", "localhost:"+port)
				exitFunc()
			case syscall.SIGUSR1:
				log.Println("usr1", s)
			case syscall.SIGUSR2:
				log.Println("usr2", s)
			default:
				log.Println("other", s)
			}
		}
	}()

	e2 := etcdAdd(cli, "hello-service", "localhost:"+port)
	if e2 != nil {
		log.Fatalf("etcdAdd failed.")
	}

	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
