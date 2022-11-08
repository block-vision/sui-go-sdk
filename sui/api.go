package sui

import (
	"github.com/block-vision/sui-go-sdk/common/rpc_client"
)

type ISuiAPI interface {
	IBaseAPI
	IWriteTransactionAPI
	IReadEventFromSuiAPI
	IReadMoveFromSuiAPI
	IReadObjectFromSuiAPI
	IReadTransactionFromSuiAPI
	IFeatureSuiAPI
}

type Client struct {
	IBaseAPI
	IWriteTransactionAPI
	IReadEventFromSuiAPI
	IReadMoveFromSuiAPI
	IReadObjectFromSuiAPI
	IReadTransactionFromSuiAPI
	IFeatureSuiAPI
}

func NewSuiClient(dest string) ISuiAPI {
	cli := rpc_client.NewRPCClient(dest)
	return &Client{
		IWriteTransactionAPI: &suiWriteTransactionImpl{
			cli: cli,
		},
		IReadEventFromSuiAPI: &suiReadEventFromSuiImpl{
			cli: cli,
		},
		IReadMoveFromSuiAPI: &suiReadMoveFromSuiImpl{
			cli: cli,
		},
		IReadObjectFromSuiAPI: &suiReadObjectFromSuiImpl{
			cli: cli,
		},
		IReadTransactionFromSuiAPI: &suiReadTransactionFromSuiImpl{
			cli: cli,
		},
		IBaseAPI:       &suiBaseImpl{cli: cli},
		IFeatureSuiAPI: &suiFeatureImpl{cli: cli},
	}
}
