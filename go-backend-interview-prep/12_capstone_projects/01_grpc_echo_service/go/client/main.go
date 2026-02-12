package main

// gRPC Echo Client
//
// Sends 3 requests to the EchoService:
//   1. Normal request
//   2. Request with custom request_id
//   3. Request with a very short deadline (demonstrates timeout)
//
// IMPORTANT: Requires generated code + running server.
// See ../generated/README.md for setup.

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-backend-interview-prep/12_capstone_projects/01_grpc_echo_service/go/generated/echopb"
)

func main() {
	// Connect to the server
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	fmt.Println("=== gRPC Echo Client ===")

	// --- Request 1: normal echo ---
	fmt.Println("\n--- Request 1: normal echo ---")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()

	resp1, err := client.Echo(ctx1, &pb.EchoRequest{
		Message:   "hello gRPC",
		RequestId: "req-001",
	})
	if err != nil {
		log.Printf("echo failed: %v", err)
	} else {
		fmt.Printf("response: message=%q request_id=%q server_time=%q\n",
			resp1.Message, resp1.RequestId, resp1.ServerTime)
	}

	// --- Request 2: different message ---
	fmt.Println("\n--- Request 2: with tracing ID ---")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	resp2, err := client.Echo(ctx2, &pb.EchoRequest{
		Message:   "interview prep",
		RequestId: "trace-abc-123",
	})
	if err != nil {
		log.Printf("echo failed: %v", err)
	} else {
		fmt.Printf("response: message=%q request_id=%q server_time=%q\n",
			resp2.Message, resp2.RequestId, resp2.ServerTime)
	}

	// --- Request 3: very short deadline (demonstrates timeout) ---
	fmt.Println("\n--- Request 3: deadline exceeded (1ms timeout) ---")
	ctx3, cancel3 := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel3()

	// Sleep briefly to ensure deadline expires before server can respond
	time.Sleep(2 * time.Millisecond)

	resp3, err := client.Echo(ctx3, &pb.EchoRequest{
		Message:   "this should timeout",
		RequestId: "req-timeout",
	})
	if err != nil {
		fmt.Printf("expected error: %v\n", err)
	} else {
		fmt.Printf("unexpected success: %v\n", resp3)
	}

	fmt.Println("\ndone")
}
