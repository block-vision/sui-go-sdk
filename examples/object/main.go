package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

var ctx = context.Background()
var cli = sui.NewSuiClient(constant.BvMainnetEndpoint)

func main() {
	SuiGetObject()
	SuiXGetOwnedObjects()
	SuiMultiGetObjects()
	SuiXGetDynamicField()
	SuiTryGetPastObject()
	SuiGetLoadedChildObjects()
	SuiTryMultiGetPastObjects()
}

func SuiGetObject() {
	rsp, err := cli.SuiGetObject(ctx, models.SuiGetObjectRequest{
		ObjectId: "0xaf647fef62f139f9bbb3ece219d40a49024331531fbe7e5ac3f5718fc1f23c62",
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)

}

func SuiXGetOwnedObjects() {
	suiObjectResponseQuery := models.SuiObjectResponseQuery{
		Filter: models.ObjectFilterByPackage{
			Package: "0x0b2041bdc2b1f9c5ff4bafc5883e6f48dd1dc81a1d4d2392d3d4c8e02bb4dd82",
		},
		Options: models.SuiObjectDataOptions{
			ShowType:                true,
			ShowContent:             true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
			ShowDisplay:             true,
		},
	}

	rsp, err := cli.SuiXGetOwnedObjects(ctx, models.SuiXGetOwnedObjectsRequest{
		Address: "0xb7f98d327f19f674347e1e40641408253142d6e7e5093a7c96eda8cdfd7d9bb5",
		Query:   suiObjectResponseQuery,
		Limit:   5,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiMultiGetObjects() {
	rsp, err := cli.SuiMultiGetObjects(ctx, models.SuiMultiGetObjectsRequest{
		ObjectIds: []string{"0x02b547f6aaece97b39142093205d5802101599833c65334ec3beb0aeed82c884"},
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, object := range rsp {
		utils.PrettyPrint(*object)
	}
}

func SuiXGetDynamicField() {
	rsp, err := cli.SuiXGetDynamicField(ctx, models.SuiXGetDynamicFieldRequest{
		ObjectId: "0x96cdd7b5b715477de928b5e8c58751ce9d26c9fe89f0d270c0d78350da9f3b4c",
		Limit:    5,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiTryGetPastObject() {
	rsp, err := cli.SuiTryGetPastObject(ctx, models.SuiTryGetPastObjectRequest{
		ObjectId: "0xeeb964d1e640219c8ddb791cc8548f3242a3392b143ff47484a3753291cad898",
		Version:  19636,
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiGetLoadedChildObjects() {
	rsp, err := cli.SuiGetLoadedChildObjects(ctx, models.SuiGetLoadedChildObjectsRequest{
		Digest: "DDvbPE1Ty138BEsu1238rRkpx4DMMDKbCJktt4H1cG6T",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}

func SuiTryMultiGetPastObjects() {
	rsp, err := cli.SuiTryMultiGetPastObjects(ctx, models.SuiTryMultiGetPastObjectsRequest{
		MultiGetPastObjects: []*models.PastObject{
			{
				ObjectId: "0xfe3e114168d65ca9c86e43ce0f8dc4f8e0fa5a03634a4c6bf292679f6d73ec72",
				Version:  "22945798",
			},
			{
				ObjectId: "0xbf67e84fef313e6f1756411b095ba07868804852c939691b300a7e1e45d0251f",
				Version:  "23119685",
			},
		},
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp)
}
