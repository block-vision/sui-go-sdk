// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// TestRequest test request structure
type TestRequest struct {
	Address string `validate:"required,checkAddress"`
	Amount  uint64 `validate:"required,gte=1"`
}

// TestResponse test response structure
type TestResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// mockServer creates a mock HTTP server
func mockServer(responseBody string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
}

// TestNewBaseRequestHandler tests the constructor
func TestNewBaseRequestHandler(t *testing.T) {
	conn := httpconn.NewHttpConn("http://test.com", nil)
	handler := NewBaseRequestHandler(conn)

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.conn)
}

// TestExecuteRequest tests basic request execution
func TestExecuteRequest(t *testing.T) {
	tests := []struct {
		name           string
		responseBody   string
		statusCode     int
		method         string
		params         []interface{}
		expectedResult TestResponse
		expectError    bool
		errorContains  string
	}{
		{
			name: "successful request with object result",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"result": {
					"success": true,
					"message": "test message",
					"data": "test data"
				}
			}`,
			statusCode: 200,
			method:     "test_method",
			params:     []interface{}{"param1", "param2"},
			expectedResult: TestResponse{
				Success: true,
				Message: "test message",
				Data:    "test data",
			},
			expectError: false,
		},
		{
			name: "error response from server",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"error": {
					"code": -32602,
					"message": "Invalid params"
				}
			}`,
			statusCode:    200,
			method:        "test_method",
			params:        []interface{}{"invalid"},
			expectError:   true,
			errorContains: "Invalid params",
		},
		{
			name: "missing result field",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1
			}`,
			statusCode:    200,
			method:        "test_method",
			params:        []interface{}{},
			expectError:   true,
			errorContains: "no result field",
		},
		{
			name: "invalid json response",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"result": "invalid json for struct"
			}`,
			statusCode:    200,
			method:        "test_method",
			params:        []interface{}{},
			expectError:   true,
			errorContains: "unmarshal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(tt.responseBody, tt.statusCode)
			defer server.Close()

			conn := httpconn.NewHttpConn(server.URL, nil)
			handler := NewBaseRequestHandler(conn)

			var result TestResponse
			err := handler.ExecuteRequest(context.Background(), tt.method, tt.params, &result)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

// TestExecuteRequestWithStringResult tests request with string result
func TestExecuteRequestWithStringResult(t *testing.T) {
	responseBody := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": "test string result"
	}`

	server := mockServer(responseBody, 200)
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	var result string
	err := handler.ExecuteRequest(context.Background(), "test_method", []interface{}{}, &result)

	assert.NoError(t, err)
	assert.Equal(t, "test string result", result)
}

// TestExecuteRequestWithUint64Result tests request with uint64 result
func TestExecuteRequestWithUint64Result(t *testing.T) {
	responseBody := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": 12345
	}`

	server := mockServer(responseBody, 200)
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	var result uint64
	err := handler.ExecuteRequest(context.Background(), "test_method", []interface{}{}, &result)

	assert.NoError(t, err)
	assert.Equal(t, uint64(12345), result)
}

// TestExecuteRequestWithValidation tests request execution with validation
func TestExecuteRequestWithValidation(t *testing.T) {
	tests := []struct {
		name          string
		request       TestRequest
		responseBody  string
		expectError   bool
		errorContains string
	}{
		{
			name: "valid request",
			request: TestRequest{
				Address: "0x0000000000000000000000000000000000000000000000000000000000000001",
				Amount:  100,
			},
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"result": {
					"success": true,
					"message": "validated",
					"data": "test"
				}
			}`,
			expectError: false,
		},
		{
			name: "invalid address",
			request: TestRequest{
				Address: "invalid_address",
				Amount:  100,
			},
			expectError:   true,
			errorContains: "validation failed",
		},
		{
			name: "missing required field",
			request: TestRequest{
				Amount: 100,
			},
			expectError:   true,
			errorContains: "validation failed",
		},
		{
			name: "amount too small",
			request: TestRequest{
				Address: "0x0000000000000000000000000000000000000000000000000000000000000001",
				Amount:  0,
			},
			expectError:   true,
			errorContains: "validation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(tt.responseBody, 200)
			defer server.Close()

			conn := httpconn.NewHttpConn(server.URL, nil)
			handler := NewBaseRequestHandler(conn)

			var result TestResponse
			err := handler.ExecuteRequestWithValidation(
				context.Background(),
				"test_method",
				[]interface{}{tt.request.Address, tt.request.Amount},
				tt.request,
				&result,
			)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				assert.True(t, result.Success)
			}
		})
	}
}

// TestExecuteSimpleRequest tests simple request execution
func TestExecuteSimpleRequest(t *testing.T) {
	tests := []struct {
		name          string
		responseBody  string
		statusCode    int
		expectError   bool
		errorContains string
		expectedValue interface{}
	}{
		{
			name: "successful simple request with string result",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"result": "simple string result"
			}`,
			statusCode:    200,
			expectError:   false,
			expectedValue: "simple string result",
		},
		{
			name: "successful simple request with number result",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"result": 42
			}`,
			statusCode:    200,
			expectError:   false,
			expectedValue: int64(42),
		},
		{
			name: "error response",
			responseBody: `{
				"jsonrpc": "2.0",
				"id": 1,
				"error": {
					"code": -32601,
					"message": "Method not found"
				}
			}`,
			statusCode:    200,
			expectError:   true,
			errorContains: "Method not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := mockServer(tt.responseBody, tt.statusCode)
			defer server.Close()

			conn := httpconn.NewHttpConn(server.URL, nil)
			handler := NewBaseRequestHandler(conn)

			result, err := handler.ExecuteSimpleRequest(
				context.Background(),
				"test_method",
				[]interface{}{"param1"},
			)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				switch expected := tt.expectedValue.(type) {
				case string:
					assert.Equal(t, expected, result.String())
				case int64:
					assert.Equal(t, expected, result.Int())
				}
			}
		})
	}
}

// TestExecuteRequestWithNetworkError tests network error scenarios
func TestExecuteRequestWithNetworkError(t *testing.T) {
	// Use invalid URL to simulate network error
	conn := httpconn.NewHttpConn("http://localhost:99999", nil) // Use non-existent port
	handler := NewBaseRequestHandler(conn)

	var result TestResponse
	err := handler.ExecuteRequest(context.Background(), "test_method", []interface{}{}, &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request failed")
}

// TestExecuteRequestWithContextCancellation tests context cancellation
func TestExecuteRequestWithContextCancellation(t *testing.T) {
	// Create a server that will delay response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate long processing time
		select {
		case <-r.Context().Done():
			return
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write([]byte(`{"jsonrpc": "2.0", "id": 1, "result": "success"}`))
		}
	}))
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	// Create a context that will be cancelled immediately
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	var result TestResponse
	err := handler.ExecuteRequest(ctx, "test_method", []interface{}{}, &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request failed")
}

// TestParseResponseEdgeCases tests edge cases of parseResponse method
func TestParseResponseEdgeCases(t *testing.T) {
	handler := &BaseRequestHandler{}

	tests := []struct {
		name          string
		respBytes     []byte
		method        string
		result        interface{}
		expectError   bool
		errorContains string
	}{
		{
			name:          "empty response",
			respBytes:     []byte(""),
			method:        "test",
			result:        &TestResponse{},
			expectError:   true,
			errorContains: "no result field",
		},
		{
			name:          "invalid json",
			respBytes:     []byte("invalid json"),
			method:        "test",
			result:        &TestResponse{},
			expectError:   true,
			errorContains: "no result field",
		},
		{
			name: "array result",
			respBytes: []byte(`{
				"jsonrpc": "2.0",
				"id": 1,
				"result": ["item1", "item2", "item3"]
			}`),
			method:      "test",
			result:      &[]string{},
			expectError: false,
		},
		{
			name: "null result",
			respBytes: []byte(`{
				"jsonrpc": "2.0",
				"id": 1,
				"result": null
			}`),
			method:      "test",
			result:      new(string), // Use string pointer to handle null value
			expectError: false,
		},
		{
			name: "boolean result to string",
			respBytes: []byte(`{
				"jsonrpc": "2.0",
				"id": 1,
				"result": true
			}`),
			method:      "test",
			result:      new(string),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.parseResponse(tt.respBytes, tt.method, tt.result)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestExecuteRequestWithArrayResponse tests array response
func TestExecuteRequestWithArrayResponse(t *testing.T) {
	responseBody := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": [
			{"success": true, "message": "msg1", "data": "data1"},
			{"success": false, "message": "msg2", "data": "data2"}
		]
	}`

	server := mockServer(responseBody, 200)
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	var result []TestResponse
	err := handler.ExecuteRequest(context.Background(), "test_method", []interface{}{}, &result)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.True(t, result[0].Success)
	assert.False(t, result[1].Success)
}

// TestExecuteRequestWithComplexParams tests complex parameters
func TestExecuteRequestWithComplexParams(t *testing.T) {
	responseBody := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": {
			"success": true,
			"message": "complex params processed",
			"data": "result"
		}
	}`

	server := mockServer(responseBody, 200)
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	// Test complex parameter types
	complexParams := []interface{}{
		"string_param",
		123,
		true,
		map[string]interface{}{
			"nested": "value",
			"number": 456,
		},
		[]string{"array", "of", "strings"},
	}

	var result TestResponse
	err := handler.ExecuteRequest(context.Background(), "test_method", complexParams, &result)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, "complex params processed", result.Message)
}

// BenchmarkExecuteRequest performance benchmark test
func BenchmarkExecuteRequest(b *testing.B) {
	responseBody := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": {
			"success": true,
			"message": "benchmark test",
			"data": "benchmark data"
		}
	}`

	server := mockServer(responseBody, 200)
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result TestResponse
		err := handler.ExecuteRequest(context.Background(), "benchmark_method", []interface{}{"param"}, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// TestConcurrentRequests tests concurrent requests
func TestConcurrentRequests(t *testing.T) {
	responseBody := `{
		"jsonrpc": "2.0",
		"id": 1,
		"result": {
			"success": true,
			"message": "concurrent test",
			"data": "concurrent data"
		}
	}`

	server := mockServer(responseBody, 200)
	defer server.Close()

	conn := httpconn.NewHttpConn(server.URL, nil)
	handler := NewBaseRequestHandler(conn)

	const numGoroutines = 10
	const requestsPerGoroutine = 5

	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines*requestsPerGoroutine)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < requestsPerGoroutine; j++ {
				var result TestResponse
				err := handler.ExecuteRequest(
					context.Background(),
					"concurrent_method",
					[]interface{}{fmt.Sprintf("param_%d_%d", goroutineID, j)},
					&result,
				)
				if err != nil {
					errors <- err
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Check for any errors
	for err := range errors {
		t.Errorf("Concurrent request failed: %v", err)
	}
}
