package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
	v2 "github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2"
)

func main() {
	fmt.Println("=== Sui gRPC Transaction Execution Service Examples ===")

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	ctx := context.Background()

	// Get transaction execution service
	txService, err := client.TransactionExecutionService(ctx)
	if err != nil {
		log.Fatalf("Failed to get transaction execution service: %v", err)
	}

	// Run example
	fmt.Println("\n1. Executing transaction...")
	exampleExecuteTransaction(ctx, txService)
}

// ExecuteTransaction - Execute a transaction on the Sui network
func exampleExecuteTransaction(ctx context.Context, service v2.TransactionExecutionServiceClient) {
	// Note: This is a simplified example. In reality, you would need to:
	// 1. Build a proper transaction using PTB (Programmable Transaction Block)
	// 2. Sign the transaction with the appropriate private key
	// 3. Provide valid transaction data and signatures

	req := &v2.ExecuteTransactionRequest{
		Transaction: &v2.Transaction{
			// Transaction data would go here
			// This should be a properly constructed transaction
		},
		Signatures: []*v2.UserSignature{
			// Valid signatures would go here
			// Each signature should correspond to a required signer
		},
		// Options for transaction execution
		// You can specify execution options like max gas budget, etc.
	}

	resp, err := service.ExecuteTransaction(ctx, req)
	if err != nil {
		fmt.Printf("❌ ExecuteTransaction failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Transaction executed successfully: %v\n", resp)
	if resp.Transaction != nil {
		fmt.Printf("   Transaction digest: %v\n", resp.Transaction.Digest)
		fmt.Printf("   Effects: %v\n", resp.Transaction.Effects)
	}
}
