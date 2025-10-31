package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
	v2 "github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2"
)

func main() {
	fmt.Println("=== Sui gRPC Name Service Examples ===")

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	ctx := context.Background()

	// Get name service
	nameService, err := client.NameService(ctx)
	if err != nil {
		log.Fatalf("Failed to get name service: %v", err)
	}

	// Run all examples
	fmt.Println("\n1. Looking up name...")
	exampleLookupName(ctx, nameService)

	fmt.Println("\n2. Reverse looking up name...")
	exampleReverseLookupName(ctx, nameService)
}

// LookupName - Look up a name to get the associated object
func exampleLookupName(ctx context.Context, service v2.NameServiceClient) {
	name := "websui.sui"
	req := &v2.LookupNameRequest{
		Name: &name,
	}

	resp, err := service.LookupName(ctx, req)
	if err != nil {
		fmt.Printf("❌ LookupName failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Name Lookup Result: %v\n", resp)
}

// ReverseLookupName - Reverse lookup to get the name associated with an address
func exampleReverseLookupName(ctx context.Context, service v2.NameServiceClient) {
	address := "0xa1c56b400b57e9fca5011c80948a0f2fadbfc407470640dc7026eae575c93c7b"
	req := &v2.ReverseLookupNameRequest{
		Address: &address,
	}

	resp, err := service.ReverseLookupName(ctx, req)
	if err != nil {
		fmt.Printf("❌ ReverseLookupName failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Reverse Lookup Result: %v\n", resp)
}
