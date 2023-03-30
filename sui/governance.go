package sui

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/shoshinsquare/sui-go-sdk/common/rpc_client"
	"github.com/shoshinsquare/sui-go-sdk/models"
	"github.com/tidwall/gjson"
)

var _ IGovernanceAPI = (*SuiGovernanceImpl)(nil)

type IGovernanceAPI interface {
	GetDelegatedStakes(ctx context.Context, req models.GetDelegatedStakesRequest, opts ...interface{}) (models.GetDelegatedStakeResponse, error)
	GetValidators(ctx context.Context, req models.GetValidatorsRequest, opts ...interface{}) (models.GetValidatorsResponse, error)
	GetCommitteeInfo(ctx context.Context, req models.GetCommitteeInfoRequest, opts ...interface{}) (models.GetCommitteeInfoResponse, error)
	GetSuiSystemState(ctx context.Context, req models.GetSuiSystemStateRequest, opt ...interface{}) (models.GetSuiSystemStateResponse, error)
	GetCheckpoint(ctx context.Context, req models.GetCheckpointRequest, opts ...interface{}) (models.GetCheckpointResponse, error)
	GetLatestCheckpointSequenceNumber(ctx context.Context, req models.GetLatestCheckpointSequenceNumberRequest, opts ...interface{}) (uint64, error)
}

type SuiGovernanceImpl struct {
	cli *rpc_client.RPCClient
}

func (s *SuiGovernanceImpl) GetCommitteeInfo(ctx context.Context, req models.GetCommitteeInfoRequest, opts ...interface{}) (models.GetCommitteeInfoResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getCommitteeInfo",
		Params: []interface{}{
			req.EpochId,
		},
	})
	if err != nil {
		return models.GetCommitteeInfoResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetCommitteeInfoResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetCommitteeInfoResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetCommitteeInfoResponse{}, err
	}
	return rsp, nil
}

func (s *SuiGovernanceImpl) GetDelegatedStakes(ctx context.Context, req models.GetDelegatedStakesRequest, opts ...interface{}) (models.GetDelegatedStakeResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getDelegatedStakes",
		Params: []interface{}{
			req.Owner,
		},
	})
	if err != nil {
		return models.GetDelegatedStakeResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetDelegatedStakeResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetDelegatedStakeResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetDelegatedStakeResponse{}, err
	}
	return rsp, nil
}

func (s *SuiGovernanceImpl) GetValidators(ctx context.Context, req models.GetValidatorsRequest, opts ...interface{}) (models.GetValidatorsResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getValidators",
		Params: []interface{}{},
	})
	if err != nil {
		return models.GetValidatorsResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetValidatorsResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetValidatorsResponse
	var arr []models.ValidatorMetadata
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &arr)
	if err != nil {
		return models.GetValidatorsResponse{}, err
	}
	rsp.Result = arr
	return rsp, nil
}

func (s *SuiGovernanceImpl) GetSuiSystemState(ctx context.Context, req models.GetSuiSystemStateRequest, opts ...interface{}) (models.GetSuiSystemStateResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getSuiSystemState",
		Params: []interface{}{},
	})
	if err != nil {
		return models.GetSuiSystemStateResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetSuiSystemStateResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetSuiSystemStateResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetSuiSystemStateResponse{}, err
	}
	return rsp, nil
}

// Get checkpoint by id
func (s *SuiGovernanceImpl) GetCheckpoint(ctx context.Context, req models.GetCheckpointRequest, opts ...interface{}) (models.GetCheckpointResponse, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getCheckpoint",
		Params: []interface{}{
			req.Id,
		},
	})
	if err != nil {
		return models.GetCheckpointResponse{}, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return models.GetCheckpointResponse{}, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp models.GetCheckpointResponse
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return models.GetCheckpointResponse{}, err
	}
	return rsp, nil
}

// Get checkpoint by id
func (s *SuiGovernanceImpl) GetLatestCheckpointSequenceNumber(ctx context.Context, req models.GetLatestCheckpointSequenceNumberRequest, opts ...interface{}) (uint64, error) {
	respBytes, err := s.cli.Request(ctx, models.Operation{
		Method: "sui_getLatestCheckpointSequenceNumber",
		Params: []interface{}{},
	})
	if err != nil {
		return 0, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return 0, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var rsp uint64
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return 0, err
	}
	return rsp, nil
}
