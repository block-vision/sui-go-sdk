package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
	v2 "github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2"
)

func main() {
	fmt.Println("=== Sui gRPC Signature Verification Service Examples ===")

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	ctx := context.Background()

	// Get signature verification service
	sigVerificationService, err := client.SignatureVerificationService(ctx)
	if err != nil {
		log.Fatalf("Failed to get signature verification service: %v", err)
	}

	// Run example
	fmt.Println("\n1. Verifying signature...")
	exampleVerifySignature(ctx, sigVerificationService)
}

// VerifySignature - Verify a signature against message
func exampleVerifySignature(ctx context.Context, service v2.SignatureVerificationServiceClient) {
	// Example signature verification request
	// Note: This is a simplified example with placeholder data
	req := &v2.VerifySignatureRequest{
		// Message and signature would be properly constructed here
		// For actual usage, you need to provide:
		// - Proper message bytes
		// - Valid UserSignature
	}

	resp, err := service.VerifySignature(ctx, req)
	if err != nil {
		fmt.Printf("❌ VerifySignature failed: %v\n", err)
		return
	}

	if resp.IsValid != nil && *resp.IsValid {
		fmt.Printf("✅ Signature is valid\n")
	} else {
		fmt.Printf("❌ Signature is invalid\n")
	}

	fmt.Printf("📋 Verification response: %v\n", resp)
}
