package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
	"github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2beta2"
)

func main() {
	fmt.Println("=== Sui gRPC Live Data Service Examples ===")

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	ctx := context.Background()

	// Get live data service
	liveDataService, err := client.LiveDataService(ctx)
	if err != nil {
		log.Fatalf("Failed to get live data service: %v", err)
	}

	// Run all examples
	fmt.Println("\n1. Listing dynamic fields...")
	exampleListDynamicFields(ctx, liveDataService)

	fmt.Println("\n2. Listing owned objects...")
	exampleListOwnedObjects(ctx, liveDataService)

	fmt.Println("\n3. Getting coin info...")
	exampleGetCoinInfo(ctx, liveDataService)

	fmt.Println("\n4. Getting balance...")
	exampleGetBalance(ctx, liveDataService)

	fmt.Println("\n5. Listing balances...")
	exampleListBalances(ctx, liveDataService)

	fmt.Println("\n6. Simulating transaction...")
	exampleSimulateTransaction(ctx, liveDataService)
}

// ListDynamicFields - List dynamic fields for an object
func exampleListDynamicFields(ctx context.Context, service v2beta2.LiveDataServiceClient) {
	parent := "0x266f5a401df5fa40fc5ab2a1a8e74ac41fe5fb241e106eb608bf37c732c17e0e"
	req := &v2beta2.ListDynamicFieldsRequest{
		// Add proper request fields based on actual proto definition
		Parent: &parent,
	}

	resp, err := service.ListDynamicFields(ctx, req)
	if err != nil {
		fmt.Printf("❌ ListDynamicFields failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Dynamic Fields response: %v\n", resp)
}

// ListOwnedObjects - List objects owned by an address
func exampleListOwnedObjects(ctx context.Context, service v2beta2.LiveDataServiceClient) {
	owner := "0xac5bceec1b789ff840d7d4e6ce4ce61c90d190a7f8c4f4ddf0bff6ee2413c33c"
	req := &v2beta2.ListOwnedObjectsRequest{
		Owner: &owner,
	}

	resp, err := service.ListOwnedObjects(ctx, req)
	if err != nil {
		fmt.Printf("❌ ListOwnedObjects failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Owned Objects response: %v\n", resp)
}

// GetCoinInfo - Get coin information
func exampleGetCoinInfo(ctx context.Context, service v2beta2.LiveDataServiceClient) {
	coinType := "0x2::sui::SUI"
	req := &v2beta2.GetCoinInfoRequest{
		CoinType: &coinType,
	}

	resp, err := service.GetCoinInfo(ctx, req)
	if err != nil {
		fmt.Printf("❌ GetCoinInfo failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Coin Info response: %v\n", resp)
}

// GetBalance - Get balance for an address and coin type
func exampleGetBalance(ctx context.Context, service v2beta2.LiveDataServiceClient) {
	owner := "0xac5bceec1b789ff840d7d4e6ce4ce61c90d190a7f8c4f4ddf0bff6ee2413c33c"
	coinType := "0x2::sui::SUI"
	req := &v2beta2.GetBalanceRequest{
		Owner:    &owner,
		CoinType: &coinType,
	}

	resp, err := service.GetBalance(ctx, req)
	if err != nil {
		fmt.Printf("❌ GetBalance failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Balance response: %v\n", resp)
}

// ListBalances - List all balances for an address
func exampleListBalances(ctx context.Context, service v2beta2.LiveDataServiceClient) {
	owner := "0xac5bceec1b789ff840d7d4e6ce4ce61c90d190a7f8c4f4ddf0bff6ee2413c33c"
	req := &v2beta2.ListBalancesRequest{
		Owner: &owner,
	}

	resp, err := service.ListBalances(ctx, req)
	if err != nil {
		fmt.Printf("❌ ListBalances failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Balances response: %v\n", resp)
}

// SimulateTransaction - Simulate a transaction
func exampleSimulateTransaction(ctx context.Context, service v2beta2.LiveDataServiceClient) {
	req := &v2beta2.SimulateTransactionRequest{
		Transaction: &v2beta2.Transaction{
			// Transaction data would go here
		},
	}

	resp, err := service.SimulateTransaction(ctx, req)
	if err != nil {
		fmt.Printf("❌ SimulateTransaction failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Simulation response: %v\n", resp)
}
