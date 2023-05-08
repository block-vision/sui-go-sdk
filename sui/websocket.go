// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"github.com/block-vision/sui-go-sdk/common/wsconn"
)

// ISuiWebsocketAPI defines the subscription API related interface, and then implement it by the WebsocketClient.
type ISuiWebsocketAPI interface {
	ISubscribeAPI
}

// WebsocketClient implements SuiWebsocketAPI related interfaces.
type WebsocketClient struct {
	ISubscribeAPI
}

// NewSuiWebsocketClient instantiates the WebsocketClient to call the methods of each module.
func NewSuiWebsocketClient(rpcUrl string) ISuiWebsocketAPI {
	conn := wsconn.NewWsConn(rpcUrl)
	return &WebsocketClient{
		ISubscribeAPI: &suiSubscribeImpl{
			conn: conn,
		},
	}
}
