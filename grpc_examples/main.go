package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/common/grpcconn"
	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
)

func main() {
	fmt.Println("=== Sui gRPC Examples Runner ===")
	fmt.Println()

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	config := utils.NewDefaultConfig()
	fmt.Printf("🔧 Configuration:\n")
	fmt.Printf("   Target: %s\n", config.Target)
	fmt.Printf("   Token: %s\n", config.Token)
	fmt.Println()

	ctx := context.Background()

	// Test connectivity to all services
	fmt.Println("🧪 Testing connectivity to all gRPC services...")
	testAllServices(ctx, client)

	fmt.Println()
	fmt.Println("📚 Available Examples:")
	fmt.Println("   1. Ledger Service      - go run grpc_examples/ledger_service/main.go")
	fmt.Println("   2. Name Service        - go run grpc_examples/name_service/main.go")
	fmt.Println("   3. State Service       - go run grpc_examples/state_service/main.go")
	fmt.Println("   4. Transaction Exec    - go run grpc_examples/transaction_execution_service/main.go")
	fmt.Println("   5. Subscription        - go run grpc_examples/subscription_service/main.go")
	fmt.Println("   6. Move Package        - go run grpc_examples/move_package_service/main.go")
	fmt.Println("   7. Signature Verify    - go run grpc_examples/signature_verification_service/main.go")
	fmt.Println()
	fmt.Println("📖 For detailed documentation, see: grpc_examples/README.md")
}

func testAllServices(ctx context.Context, client *grpcconn.SuiGrpcClient) {
	services := []struct {
		name     string
		testFunc func(context.Context, *grpcconn.SuiGrpcClient) error
	}{
		{"LedgerService", testLedgerService},
		{"NameService", testNameService},
		{"StateService", testStateService},
		{"MovePackageService", testMovePackageService},
		{"SubscriptionService", testSubscriptionService},
		{"TransactionExecutionService", testTransactionExecutionService},
		{"SignatureVerificationService", testSignatureVerificationService},
	}

	for _, service := range services {
		fmt.Printf("   Testing %s... ", service.name)
		err := service.testFunc(ctx, client)
		if err != nil {
			fmt.Printf("❌ Failed: %v\n", err)
		} else {
			fmt.Printf("✅ Connected\n")
		}
	}
}

func testLedgerService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.LedgerService(ctx)
	return err
}

func testNameService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.NameService(ctx)
	return err
}

func testStateService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.StateService(ctx)
	return err
}

func testMovePackageService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.MovePackageService(ctx)
	return err
}

func testSubscriptionService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.SubscriptionService(ctx)
	return err
}

func testTransactionExecutionService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.TransactionExecutionService(ctx)
	return err
}

func testSignatureVerificationService(ctx context.Context, client *grpcconn.SuiGrpcClient) error {
	_, err := client.SignatureVerificationService(ctx)
	return err
}
