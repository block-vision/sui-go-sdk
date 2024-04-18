package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/block-vision/sui-go-sdk/common/httpconn"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

type IReadNameServiceFromSuiAPI interface {
	SuiXResolveNameServiceAddress(ctx context.Context, req models.SuiXResolveNameServiceAddressRequest) (string, error)
	SuiXResolveNameServiceNames(ctx context.Context, req models.SuiXResolveNameServiceNamesRequest) (models.SuiXResolveNameServiceNamesResponse, error)
}

type suiReadNameServiceFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiXResolveNameServiceAddress implements the method `suix_resolveNameServiceAddress`, get the resolved address given resolver and name.
func (s *suiReadNameServiceFromSuiImpl) SuiXResolveNameServiceAddress(ctx context.Context, req models.SuiXResolveNameServiceAddressRequest) (string, error) {
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_resolveNameServiceAddress",
		Params: []interface{}{
			req.Name,
		},
	})
	if err != nil {
		return "", err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return "", errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}

	return gjson.ParseBytes(respBytes).Get("result").String(), nil
}

// SuiXResolveNameServiceNames implements the method `suix_resolveNameServiceNames`, return the resolved names given address, if multiple names are resolved, the first one is the primary name.
func (s *suiReadNameServiceFromSuiImpl) SuiXResolveNameServiceNames(ctx context.Context, req models.SuiXResolveNameServiceNamesRequest) (models.SuiXResolveNameServiceNamesResponse, error) {
	var rsp models.SuiXResolveNameServiceNamesResponse
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
		Method: "suix_resolveNameServiceNames",
		Params: []interface{}{
			req.Address,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
