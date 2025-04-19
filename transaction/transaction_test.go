package transaction

import (
	"testing"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/google/go-cmp/cmp"
	"github.com/mr-tron/base58"
)

func TestNewTransaction(t *testing.T) {
	cases := []struct {
		name            string
		fun             func() *Transaction
		expectBcsBase64 string
	}{
		{
			name: "basic tx",
			fun: func() *Transaction {
				return setupTransaction()
			},
			expectBcsBase64: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAWFiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiAgAAAAAAAAAgAAECAwQFBgcICQABAgMEBQYHCAkAAQIDBAUGBwgJAQIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgUAAAAAAAAAZAAAAAAAAAAA",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tx := c.fun()

			bcs, err := tx.build(true)
			if err != nil {
				t.Fatalf("failed to marshal transaction: %v", err)
			}

			if diff := cmp.Diff(bcs, c.expectBcsBase64); diff != "" {
				t.Errorf("Transaction mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func generateObjectRef() SuiObjectRef {
	objectId := "0x6162636162636162636162636162636162636162636162636162636162636162"

	bytes := []byte{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2,
	}

	digest := base58.Encode(bytes)

	objectIdBytes, _ := ConvertSuiAddressStringToBytes(models.SuiAddress(objectId))
	digestBytes, _ := ConvertObjectDigestStringToBytes(models.ObjectDigest(digest))

	return SuiObjectRef{
		ObjectId: *objectIdBytes,
		Version:  2,
		Digest:   *digestBytes,
	}
}

func setupTransaction() *Transaction {
	tx := NewTransaction()
	tx.SetSender("0x2").
		SetGasPrice(5).SetGasBudget(100).
		SetGasPayment([]SuiObjectRef{generateObjectRef()})
	return tx
}
