package main

import (
	_ "cx-sdk/pb/com/coralogix/alerts/v2/alert_condition.proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var address = "localhost:50051"

func main() {
	fmt.Println("Hello, World!")
	conn, err := grpc.NewClient(address, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return
	}
	defer conn.Close()
}
