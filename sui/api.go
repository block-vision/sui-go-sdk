package sui

import (
	"github.com/block-vision/sui-go-sdk/httpconn"
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
	conn := httpconn.NewHttpConn(dest)
	return &Client{
		IWriteTransactionAPI: &suiWriteTransactionImpl{
			conn: conn,
		},
		IReadEventFromSuiAPI: &suiReadEventFromSuiImpl{
			conn: conn,
		},
		IReadMoveFromSuiAPI: &suiReadMoveFromSuiImpl{
			conn: conn,
		},
		IReadObjectFromSuiAPI: &suiReadObjectFromSuiImpl{
			conn: conn,
		},
		IReadTransactionFromSuiAPI: &suiReadTransactionFromSuiImpl{
			conn: conn,
		},
		IBaseAPI:       &suiBaseImpl{conn: conn},
		IFeatureSuiAPI: &suiFeatureImpl{conn: conn},
	}
}
