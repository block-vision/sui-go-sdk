# Sui gRPC Examples

This directory contains comprehensive examples for all Sui gRPC services. Each service has its own directory with specific examples demonstrating all available methods.

## Available Services

### 1. [Ledger Service](./ledger_service/) 
Query blockchain data and state information.

**Methods:**
- `GetServiceInfo` - Get service information
- `GetObject` - Get object by ID
- `BatchGetObjects` - Get multiple objects
- `GetTransaction` - Get transaction by digest
- `BatchGetTransactions` - Get multiple transactions
- `GetCheckpoint` - Get checkpoint by sequence number
- `GetEpoch` - Get epoch information

### 2. [Name Service](./name_service/)
Resolve and manage Sui names.

**Methods:**
- `LookupName` - Resolve name to object
- `ReverseLookupName` - Get name from object ID

### 3. [Transaction Execution Service](./transaction_execution_service/)
Execute transactions on the Sui network.

**Methods:**
- `ExecuteTransaction` - Submit and execute transactions

### 4. [Subscription Service](./subscription_service/)
Subscribe to real-time blockchain events.

**Methods:**
- `SubscribeCheckpoints` - Stream checkpoint updates

### 5. [Live Data Service](./live_data_service/)
Access live blockchain data.

**Methods:**
- `GetObject` - Get live object data
- `GetTransaction` - Get live transaction data
- `QueryEvents` - Query live events

### 6. [Move Package Service](./move_package_service/)
Interact with Move packages and modules.

**Methods:**
- `GetPackage` - Get package information
- `GetModule` - Get module details
- `GetFunction` - Get function information
- `GetStruct` - Get struct definition

### 7. [Signature Verification Service](./signature_verification_service/)
Verify signatures and cryptographic operations.

**Methods:**
- `VerifySignature` - Verify message signatures
- `VerifyTransactionSignature` - Verify transaction signatures

## Quick Start

1. **Update Configuration**: Each example uses placeholder values. Update these in each main.go file:
   ```go
   target := "your-grpc-server:9000"  // Your Sui gRPC endpoint
   token := "your-api-token"          // Your authentication token
   ```

2. **Run Individual Examples**:
   ```bash
   # Run ledger service examples
   go run grpc_examples/ledger_service/main.go
   
   # Run name service examples
   go run grpc_examples/name_service/main.go
   
   # Run transaction execution examples
   go run grpc_examples/transaction_execution_service/main.go
   ```

3. **Build All Examples**:
   ```bash
   # Build all examples to check for compilation errors
   go build ./grpc_examples/...
   ```

## Configuration

All examples support the following configuration options:

```go
opts := []grpcconn.GrpcConnOption{
    grpcconn.WithTimeout(time.Second * 30),     // Request timeout
    grpcconn.WithRetryCount(3),                 // Retry attempts
    grpcconn.WithDialOptions(                   // Custom gRPC options
        grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
    ),
}
```

## Authentication

The examples use the `NewSuiGrpcClientWithAuth` function which automatically adds authentication headers to all requests. Supported authentication formats:

- `Authorization: Bearer {token}`
- `x-api-key: {token}`
- `x-token: {token}`

## Error Handling

All examples include proper error handling and demonstrate:

- Connection error recovery
- Request timeout handling
- Retry mechanisms
- Graceful failure modes

## Real Data vs Examples

The examples use placeholder data for demonstration. In real usage:

1. **Object IDs**: Use actual Sui object IDs from your network
2. **Transaction Digests**: Use real transaction hashes
3. **Addresses**: Use valid Sui addresses
4. **Signatures**: Provide properly signed data
5. **Timestamps**: Use appropriate time ranges

## Streaming Services

The Subscription Service demonstrates gRPC streaming:

- **Server Streaming**: Continuous data flow from server
- **Error Recovery**: Automatic reconnection on stream errors
- **Backpressure**: Proper handling of high-volume streams
- **Graceful Shutdown**: Clean stream termination

## Best Practices

1. **Connection Management**: Reuse gRPC clients across requests
2. **Context Handling**: Use appropriate timeouts and cancellation
3. **Error Recovery**: Implement retry logic for transient failures
4. **Resource Cleanup**: Always close clients and streams
5. **Rate Limiting**: Respect service rate limits
6. **Field Masks**: Use field masks to optimize response size

## Troubleshooting

**Connection Issues:**
- Verify server address and port
- Check network connectivity
- Ensure gRPC service is enabled on the server

**Authentication Issues:**
- Verify API token is valid
- Check token format and headers
- Ensure proper TLS configuration

**Request Failures:**
- Validate request parameters
- Check field mask syntax
- Verify object IDs and addresses exist

## Related Documentation

- [gRPC Client SDK](../common/grpcconn/README.md)
- [Protocol Buffers](../pb/sui/rpc/v2beta2/)
- [Sui gRPC API Documentation](https://docs.sui.io/)

## Contributing

When adding new examples:

1. Follow the existing structure and naming conventions
2. Include comprehensive error handling
3. Add meaningful comments and documentation
4. Test with real Sui network data
5. Update this README with new service information
