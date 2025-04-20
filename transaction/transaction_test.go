package transaction

import (
	"fmt"
	"testing"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/google/go-cmp/cmp"
	"github.com/mr-tron/base58"
	"github.com/samber/lo"
)

func TestNewTransaction(t *testing.T) {
	cases := []struct {
		name                string
		fun                 func() *Transaction
		onlyTransactionKind bool
		expectBcsBase64     string
	}{
		{
			name: "tx only kind",
			fun: func() *Transaction {
				return setupTransaction()
			},
			onlyTransactionKind: true,
			expectBcsBase64:     "AAAA",
		},
		{
			name: "tx setup",
			fun: func() *Transaction {
				tx := setupTransaction()
				return tx
			},
			onlyTransactionKind: false,
			expectBcsBase64:     "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAWFiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiAgAAAAAAAAAgAAECAwQFBgcICQABAgMEBQYHCAkAAQIDBAUGBwgJAQIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgUAAAAAAAAAZAAAAAAAAAAA",
		},
		{
			name: "tx with expiration",
			fun: func() *Transaction {
				tx := setupTransaction()
				tx.SetExpiration(TransactionExpiration{
					Epoch: lo.ToPtr(uint64(100)),
				})
				return tx
			},
			onlyTransactionKind: false,
			expectBcsBase64:     "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAWFiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiY2FiAgAAAAAAAAAgAAECAwQFBgcICQABAgMEBQYHCAkAAQIDBAUGBwgJAQIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgUAAAAAAAAAZAAAAAAAAAABZAAAAAAAAAA=",
		},
		{
			name: "tx transfer",
			fun: func() *Transaction {
				tx := setupTransaction()
				splitCoin := tx.SplitCoins(tx.Gas(), []Argument{
					tx.Pure(uint64(1000000000 * 0.1)),
				})
				tx.transferObjects([]Argument{splitCoin}, tx.Pure("0x9"))
				return tx
			},
			onlyTransactionKind: false,
			expectBcsBase64:     "AAACAAgA4fUFAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAkCAgABAQAAAQECAAABAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgFhYmNhYmNhYmNhYmNhYmNhYmNhYmNhYmNhYmNhYmNhYgIAAAAAAAAAIAABAgMEBQYHCAkAAQIDBAUGBwgJAAECAwQFBgcICQECAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAYFAAAAAAAAAGQAAAAAAAAAAA==",
		},
	}

	for _, c := range cases {
		fmt.Println("Starting test case:", c.name)
		t.Run(c.name, func(t *testing.T) {
			tx := c.fun()

			bcs, err := tx.build(c.onlyTransactionKind)
			if err != nil {
				t.Fatalf("failed to marshal transaction: %v", err)
			}

			if diff := cmp.Diff(c.expectBcsBase64, bcs); diff != "" {
				t.Errorf("Transaction mismatch (-want +got):\n%s", diff)
			}

			fmt.Println(bcs)
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
		SetGasPrice(5).
		SetGasBudget(100).
		SetGasPayment([]SuiObjectRef{generateObjectRef()}).
		SetGasOwner("0x6")
	return tx
}
