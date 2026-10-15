// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package grpc_client

import (
	"reflect"
	"testing"
	"time"

	v2proto "github.com/block-vision/sui-go-sdk/pb/sui/rpc/v2"
	. "github.com/block-vision/sui-go-sdk/sui/v2/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvertGrpcInputFundsWithdrawal(t *testing.T) {
	got := convertGrpcInput(&v2proto.Input{
		Kind: v2proto.Input_FUNDS_WITHDRAWAL.Enum(),
		FundsWithdrawal: &v2proto.FundsWithdrawal{
			Amount:   proto.Uint64(6876396),
			CoinType: proto.String("0xdba3::usdc::USDC"),
			Source:   v2proto.FundsWithdrawal_SENDER.Enum(),
		},
	})
	want := &FundsWithdrawalInput{Amount: "6876396", CoinType: "0xdba3::usdc::USDC", Source: FundsWithdrawalSourceSender}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

// Inputs of unknown kinds must keep their slot so InputArg indices stay aligned.
func TestConvertGrpcTransactionDataKeepsInputIndices(t *testing.T) {
	tx := &v2proto.Transaction{
		Kind: &v2proto.TransactionKind{
			Data: &v2proto.TransactionKind_ProgrammableTransaction{ProgrammableTransaction: &v2proto.ProgrammableTransaction{
				Inputs: []*v2proto.Input{
					{Kind: v2proto.Input_PURE.Enum(), Pure: []byte{1}},
					{Kind: v2proto.Input_FUNDS_WITHDRAWAL.Enum(), FundsWithdrawal: &v2proto.FundsWithdrawal{}},
					{Kind: v2proto.Input_InputKind(99).Enum()},
					{Kind: v2proto.Input_PURE.Enum(), Pure: []byte{2}},
				},
			}},
		},
	}
	inputs := convertGrpcTransactionData(tx).Inputs
	if len(inputs) != 4 {
		t.Fatalf("got %d inputs, want 4", len(inputs))
	}
	if _, ok := inputs[1].(*FundsWithdrawalInput); !ok {
		t.Errorf("inputs[1] = %T, want *FundsWithdrawalInput", inputs[1])
	}
	if _, ok := inputs[2].(RawInput); !ok {
		t.Errorf("inputs[2] = %T, want RawInput", inputs[2])
	}
	if p, ok := inputs[3].(*PureBCSInput); !ok || p.Bytes != "Ag==" {
		t.Errorf("inputs[3] = %#v, want the second pure input", inputs[3])
	}
}

func TestConvertGrpcValidityExpiration(t *testing.T) {
	ts := time.UnixMilli(1792031436733)
	tx := &v2proto.Transaction{Expiration: &v2proto.TransactionExpiration{
		Kind:             v2proto.TransactionExpiration_VALIDITY.Enum(),
		Epoch:            proto.Uint64(1001),
		MinEpoch:         proto.Uint64(1000),
		MaxTimestamp:     timestamppb.New(ts),
		Chain:            proto.String("4btiuiMPvEENsttpZC7CZ53DruC3MAgfznDbASZ7DR6S"),
		Nonce:            proto.Uint32(7),
		AllowedProposers: &v2proto.AllowedProposers{Epoch: proto.Uint64(1000), Proposers: []uint32{1, 5}},
	}}
	minEpoch, maxEpoch, maxTs := "1000", "1001", "1792031436733"
	want := &ValidDuringExpiration{
		MinEpoch:         &minEpoch,
		MaxEpoch:         &maxEpoch,
		MaxTimestamp:     &maxTs,
		Chain:            "4btiuiMPvEENsttpZC7CZ53DruC3MAgfznDbASZ7DR6S",
		Nonce:            7,
		AllowedProposers: &AllowedProposers{Epoch: "1000", Proposers: []uint32{1, 5}},
	}
	if got := convertGrpcTransactionData(tx).Expiration; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

func TestGrpcOutputStateAccumulatorWrite(t *testing.T) {
	if got := grpcOutputState(v2proto.ChangedObject_OUTPUT_OBJECT_STATE_ACCUMULATOR_WRITE); got != ChangedObjectOutputAccumulatorWriteV1 {
		t.Fatalf("got %q", got)
	}
}

func TestConvertGrpcBalance(t *testing.T) {
	got := convertGrpcBalance(&v2proto.Balance{
		CoinType:       proto.String("0x2::sui::SUI"),
		Balance:        proto.Uint64(150),
		AddressBalance: proto.Uint64(100),
		CoinBalance:    proto.Uint64(50),
	})
	want := Balance{CoinType: "0x2::sui::SUI", Balance_: "150", CoinBalance: "50", AddressBalance: "100"}
	if got != want {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
