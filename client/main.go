package main

import (
	"context"
	"fmt"
	math "grpc/server/math"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	client := math.NewMathUtilClient(conn)
	resp, err := client.Sub(context.Background(), &math.Vector2{Numa: 6, Numb: 2})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.GetNum())

}
