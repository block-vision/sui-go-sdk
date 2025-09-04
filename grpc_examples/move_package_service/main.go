package main

import (
	"context"
	"fmt"
	"log"

	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
	"github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2beta2"
)

func main() {
	fmt.Println("=== Sui gRPC Move Package Service Examples ===")

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	ctx := context.Background()

	// Get move package service
	movePackageService, err := client.MovePackageService(ctx)
	if err != nil {
		log.Fatalf("Failed to get move package service: %v", err)
	}

	// Run all examples
	fmt.Println("\n1. Getting package...")
	exampleGetPackage(ctx, movePackageService)

	fmt.Println("\n2. Getting datatype...")
	exampleGetDatatype(ctx, movePackageService)

	fmt.Println("\n3. Getting function...")
	exampleGetFunction(ctx, movePackageService)

	fmt.Println("\n4. Listing package versions...")
	exampleListPackageVersions(ctx, movePackageService)
}

// GetPackage - Get Move package information
func exampleGetPackage(ctx context.Context, service v2beta2.MovePackageServiceClient) {
	packageId := "0x0000000000000000000000000000000000000000000000000000000000000002" // Standard library
	req := &v2beta2.GetPackageRequest{
		PackageId: &packageId,
	}

	resp, err := service.GetPackage(ctx, req)
	if err != nil {
		fmt.Printf("❌ GetPackage failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Package: %v\n", resp.Package)
}

// GetDatatype - Get Move datatype information
func exampleGetDatatype(ctx context.Context, service v2beta2.MovePackageServiceClient) {
	packageId := "0x0000000000000000000000000000000000000000000000000000000000000002"
	moduleName := "coin"
	datatypeName := "Coin"
	req := &v2beta2.GetDatatypeRequest{
		PackageId:  &packageId,
		ModuleName: &moduleName,
		Name:       &datatypeName,
	}

	resp, err := service.GetDatatype(ctx, req)
	if err != nil {
		fmt.Printf("❌ GetDatatype failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Datatype: %v\n", resp.Datatype)
}

// GetFunction - Get Move function information
func exampleGetFunction(ctx context.Context, service v2beta2.MovePackageServiceClient) {
	packageId := "0x0000000000000000000000000000000000000000000000000000000000000002"
	moduleName := "coin"
	functionName := "supply_mut"
	req := &v2beta2.GetFunctionRequest{
		PackageId:  &packageId,
		ModuleName: &moduleName,
		Name:       &functionName,
	}

	resp, err := service.GetFunction(ctx, req)
	if err != nil {
		fmt.Printf("❌ GetFunction failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Function: %v\n", resp.Function)
}

// ListPackageVersions - List versions of a package
func exampleListPackageVersions(ctx context.Context, service v2beta2.MovePackageServiceClient) {
	packageId := "0x0000000000000000000000000000000000000000000000000000000000000002"
	req := &v2beta2.ListPackageVersionsRequest{
		PackageId: &packageId,
	}

	resp, err := service.ListPackageVersions(ctx, req)
	if err != nil {
		fmt.Printf("❌ ListPackageVersions failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Package Versions: Found %d versions\n", len(resp.Versions))
}
