package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/block-vision/sui-go-sdk/grpc_examples/utils"
	"github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2beta2"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	fmt.Println("=== Sui gRPC Subscription Service Examples ===")

	// Create authenticated gRPC client using common utility
	client := utils.CreateGrpcClientWithDefaults()
	defer client.Close()

	ctx := context.Background()

	// Get subscription service
	subscriptionService, err := client.SubscriptionService(ctx)
	if err != nil {
		log.Fatalf("Failed to get subscription service: %v", err)
	}

	// Run example
	fmt.Println("\n1. Subscribing to checkpoints...")
	exampleSubscribeCheckpoints(ctx, subscriptionService)
}

// SubscribeCheckpoints - Subscribe to the stream of checkpoints
func exampleSubscribeCheckpoints(ctx context.Context, service v2beta2.SubscriptionServiceClient) {
	// Create context with timeout for the streaming operation
	// streamCtx, cancel := context.WithTimeout(ctx, time.Minute*5)
	// defer cancel()

	req := &v2beta2.SubscribeCheckpointsRequest{
		ReadMask: &fieldmaskpb.FieldMask{
			Paths: []string{"*"}, // Get all fields
		},
	}

	stream, err := service.SubscribeCheckpoints(ctx, req)
	if err != nil {
		fmt.Printf("❌ SubscribeCheckpoints failed to start: %v\n", err)
		return
	}

	fmt.Println("✅ Checkpoint subscription started")
	fmt.Println("📡 Listening for checkpoint updates...")

	// Listen for checkpoint updates
	checkpointCount := 0
	maxCheckpoints := 10 // Limit for demo purposes
	for checkpointCount < maxCheckpoints {
		checkpoint, err := stream.Recv()
		if err != nil {
			fmt.Printf("❌ Error receiving checkpoint: %v\n", err)
			break
		}

		checkpointCount++
		fmt.Printf("📦 Checkpoint %d received:\n", checkpointCount)

		if checkpoint.Checkpoint != nil {
			fmt.Printf("   Sequence Number: %v\n", checkpoint.Checkpoint.SequenceNumber)
			fmt.Printf("   Digest: %v\n", checkpoint.Checkpoint.Digest)
			fmt.Printf("   Summary: %v\n", checkpoint.Checkpoint.Summary)
		}

		// Add a small delay to make the output readable
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("✅ Received %d checkpoints\n", checkpointCount)
}

// Example of handling subscription with error recovery
func exampleSubscribeCheckpointsWithRetry(ctx context.Context, service v2beta2.SubscriptionServiceClient) {
	fmt.Println("\n2. Subscribing to checkpoints with retry logic...")

	maxRetries := 3
	retryDelay := time.Second * 5

	for attempt := 0; attempt < maxRetries; attempt++ {
		fmt.Printf("🔄 Subscription attempt %d/%d\n", attempt+1, maxRetries)

		streamCtx, cancel := context.WithTimeout(ctx, time.Minute*2)
		defer cancel()

		req := &v2beta2.SubscribeCheckpointsRequest{
			// Start from latest if this is a retry
		}

		stream, err := service.SubscribeCheckpoints(streamCtx, req)
		if err != nil {
			fmt.Printf("❌ Subscription attempt %d failed: %v\n", attempt+1, err)
			if attempt < maxRetries-1 {
				fmt.Printf("⏳ Retrying in %v...\n", retryDelay)
				time.Sleep(retryDelay)
				continue
			}
			return
		}

		fmt.Println("✅ Subscription established")

		// Handle streaming data
		for {
			checkpoint, err := stream.Recv()
			if err != nil {
				fmt.Printf("❌ Stream error: %v\n", err)
				break // Break inner loop to retry
			}

			// Process checkpoint
			if checkpoint.Checkpoint != nil {
				fmt.Printf("📦 Checkpoint %v received\n", checkpoint.Checkpoint.SequenceNumber)
			}

			// For demo purposes, exit after receiving a few checkpoints
			// In real applications, this would run continuously
			return
		}

		if attempt < maxRetries-1 {
			fmt.Printf("⏳ Retrying subscription in %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	fmt.Printf("❌ Failed to establish stable subscription after %d attempts\n", maxRetries)
}

// Example of subscription with filtering (if supported)
func exampleSubscribeCheckpointsWithFilter(ctx context.Context, service v2beta2.SubscriptionServiceClient) {
	fmt.Println("\n3. Subscribing to checkpoints with filtering...")

	streamCtx, cancel := context.WithTimeout(ctx, time.Minute*2)
	defer cancel()

	req := &v2beta2.SubscribeCheckpointsRequest{
		// Add any filtering options if supported by the service
		// This would depend on the specific proto definition
	}

	stream, err := service.SubscribeCheckpoints(streamCtx, req)
	if err != nil {
		fmt.Printf("❌ Filtered subscription failed: %v\n", err)
		return
	}

	fmt.Println("✅ Filtered checkpoint subscription started")

	// Process filtered checkpoints
	for i := 0; i < 5; i++ { // Limit for demo
		checkpoint, err := stream.Recv()
		if err != nil {
			fmt.Printf("❌ Error in filtered stream: %v\n", err)
			break
		}

		if checkpoint.Checkpoint != nil {
			fmt.Printf("🔍 Filtered checkpoint %v\n", checkpoint.Checkpoint.SequenceNumber)
		}
	}
}
