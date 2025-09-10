package object

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/transaction"
)

// ObjectReader defines the interface for reading Sui objects.
type ObjectReader interface {
	SuiGetObject(ctx context.Context, req models.SuiGetObjectRequest) (models.SuiObjectResponse, error)
}

// GetSharedObjectRef get shared object reference from sui object id
func GetSharedObjectRef(ctx context.Context, client ObjectReader, objectId string, mutable bool) (*transaction.SharedObjectRef, error) {
	rsp, err := client.SuiGetObject(ctx, models.SuiGetObjectRequest{ObjectId: objectId, Options: models.SuiObjectDataOptions{
		ShowBcs:                 true,
		ShowOwner:               true,
		ShowPreviousTransaction: true,
		ShowDisplay:             true,
		ShowType:                true,
		ShowContent:             true,
		ShowStorageRebate:       true,
	}})
	if err != nil {
		return nil, err
	}
	if value, ok := rsp.Data.Owner.(map[string]any)["Shared"]; ok {
		rv := value.(map[string]interface{})["initial_shared_version"].(float64)
		obj, _ := transaction.ConvertSuiAddressStringToBytes(models.SuiAddress(objectId))
		sharedObj := transaction.SharedObjectRef{
			ObjectId:             *obj,
			InitialSharedVersion: uint64(rv),
			Mutable:              mutable,
		}
		return &sharedObj, nil
	} else {
		return nil, fmt.Errorf("object is not a shared object")
	}
}

// GetSuiObjectRef get sui object reference from sui object id
func GetSuiObjectRef(ctx context.Context, client ObjectReader, objectId string) (*transaction.SuiObjectRef, error) {
	rsp, err := client.SuiGetObject(ctx, models.SuiGetObjectRequest{ObjectId: objectId, Options: models.SuiObjectDataOptions{
		ShowBcs:                 true,
		ShowOwner:               true,
		ShowPreviousTransaction: true,
		ShowDisplay:             true,
		ShowType:                true,
		ShowContent:             true,
		ShowStorageRebate:       true,
	}})
	if err != nil {
		return nil, err
	}
	if _, ok := rsp.Data.Owner.(map[string]any)["Shared"]; ok {
		return nil, fmt.Errorf("object is a shared object")
	} else {
		obj, err := transaction.NewSuiObjectRef(models.SuiAddress(objectId), rsp.Data.Version, models.ObjectDigest(rsp.Data.Digest))
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
}
