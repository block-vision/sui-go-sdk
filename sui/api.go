// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"net/http"

	"github.com/block-vision/sui-go-sdk/common/httpconn"
)

// ISuiAPI defines the SuiAPI related interface, and then implement it by the client.
type ISuiAPI interface {
	IBaseAPI
	IReadCoinFromSuiAPI
	IWriteTransactionAPI
	IReadEventFromSuiAPI
	IReadObjectFromSuiAPI
	IReadTransactionFromSuiAPI
	IReadSystemFromSuiAPI
	IReadMoveFromSuiAPI
	IReadNameServiceFromSuiAPI
}

// Client implements SuiAPI related interfaces.
type Client struct {
	IBaseAPI
	IReadCoinFromSuiAPI
	IWriteTransactionAPI
	IReadEventFromSuiAPI
	IReadObjectFromSuiAPI
	IReadTransactionFromSuiAPI
	IReadSystemFromSuiAPI
	IReadMoveFromSuiAPI
	IReadNameServiceFromSuiAPI
}

// NewSuiClient instantiates the Sui client to call the methods of each module.
func NewSuiClient(rpcUrl string) ISuiAPI {
	conn := httpconn.NewHttpConn(rpcUrl)
	return newClient(conn)
}

// NewSuiClientWithCustomClient custom HTTP client, instantiates the Sui client to call the methods of each module.
func NewSuiClientWithCustomClient(rpcUrl string, c *http.Client) ISuiAPI {
	conn := httpconn.NewCustomHttpConn(rpcUrl, c)
	return newClient(conn)
}

// newClient return the Sui client to call the methods of each module.
func newClient(conn *httpconn.HttpConn) *Client {
	return &Client{
		IBaseAPI: &suiBaseImpl{
			conn: conn,
		},
		IReadCoinFromSuiAPI: &suiReadCoinFromSuiImpl{
			conn: conn,
		},
		IWriteTransactionAPI: &suiWriteTransactionImpl{
			conn: conn,
		},
		IReadEventFromSuiAPI: &suiReadEventFromSuiImpl{
			conn: conn,
		},
		IReadObjectFromSuiAPI: &suiReadObjectFromSuiImpl{
			conn: conn,
		},
		IReadTransactionFromSuiAPI: &suiReadTransactionFromSuiImpl{
			conn: conn,
		},
		IReadSystemFromSuiAPI: &suiReadSystemFromSuiImpl{
			conn: conn,
		},
		IReadMoveFromSuiAPI: &suiReadMoveFromSuiImpl{
			conn: conn,
		},
		IReadNameServiceFromSuiAPI: &suiReadNameServiceFromSuiImpl{
			conn: conn,
		},
	}
}
