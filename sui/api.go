package sui

import (
	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
)

type ISuiAPI interface {
	IBaseAPI
	IWriteTransactionAPI
	IReadEventFromSuiAPI
	IReadMoveFromSuiAPI
	IReadObjectFromSuiAPI
	IReadTransactionFromSuiAPI
	IFeatureSuiAPI
	ICoinAPI
	IGovernanceAPI
}

type Client struct {
	IBaseAPI
	IWriteTransactionAPI
	IReadEventFromSuiAPI
	IReadMoveFromSuiAPI
	IReadObjectFromSuiAPI
	IReadTransactionFromSuiAPI
	IFeatureSuiAPI
	ICoinAPI
	IGovernanceAPI
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
		ICoinAPI: &SuiCoinImpl{
			cli: cli,
		},
		IGovernanceAPI: &SuiGovernanceImpl{cli: cli},
		IBaseAPI:       &suiBaseImpl{cli: cli},
		IFeatureSuiAPI: &suiFeatureImpl{cli: cli},
	}
}
