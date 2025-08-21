// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/tidwall/gjson"
)

// BaseRequestHandler unified request handler to eliminate duplicate code in all API files
type BaseRequestHandler struct {
	conn *httpconn.HttpConn
}

// NewBaseRequestHandler creates a new base request handler
func NewBaseRequestHandler(conn *httpconn.HttpConn) *BaseRequestHandler {
	return &BaseRequestHandler{
		conn: conn,
	}
}

// ExecuteRequest unified request execution method
func (h *BaseRequestHandler) ExecuteRequest(ctx context.Context, method string, params []interface{}, result interface{}) error {
	respBytes, err := h.conn.Request(ctx, httpconn.Operation{
		Method: method,
		Params: params,
	})
	if err != nil {
		return fmt.Errorf("%s request failed: %w", method, err)
	}

	return h.parseResponse(respBytes, method, result)
}

// ExecuteRequestWithValidation request execution method with validation
func (h *BaseRequestHandler) ExecuteRequestWithValidation(ctx context.Context, method string, params []interface{}, req interface{}, result interface{}) error {
	if err := validate.ValidateStruct(req); err != nil {
		return fmt.Errorf("validation failed for %s: %w", method, err)
	}
	return h.ExecuteRequest(ctx, method, params, result)
}

// ExecuteSimpleRequest executes simple requests (returns basic types)
func (h *BaseRequestHandler) ExecuteSimpleRequest(ctx context.Context, method string, params []interface{}) (gjson.Result, error) {
	respBytes, err := h.conn.Request(ctx, httpconn.Operation{
		Method: method,
		Params: params,
	})
	if err != nil {
		return gjson.Result{}, fmt.Errorf("%s request failed: %w", method, err)
	}

	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return gjson.Result{}, errors.New(parsedJson.Get("error").String())
	}

	return parsedJson.Get("result"), nil
}

// parseResponse unified response parsing method
func (h *BaseRequestHandler) parseResponse(respBytes []byte, method string, result interface{}) error {
	parsedJson := gjson.ParseBytes(respBytes)
	if parsedJson.Get("error").Exists() {
		return errors.New(parsedJson.Get("error").String())
	}

	resultData := parsedJson.Get("result")
	if !resultData.Exists() {
		return fmt.Errorf("no result field in %s response", method)
	}

	// Direct return for string types
	if strPtr, ok := result.(*string); ok {
		*strPtr = resultData.String()
		return nil
	}

	// Direct return for uint64 types
	if uint64Ptr, ok := result.(*uint64); ok {
		*uint64Ptr = resultData.Uint()
		return nil
	}

	// Use JSON deserialization for complex types
	var jsonData string
	if resultData.IsObject() || resultData.IsArray() {
		jsonData = resultData.Raw
	} else {
		jsonData = resultData.String()
	}

	err := json.Unmarshal([]byte(jsonData), result)
	if err != nil {
		return fmt.Errorf("unmarshal %s response error: %w, response: %s", method, err, string(respBytes))
	}

	return nil
}
