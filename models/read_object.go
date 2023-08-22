package models

type SuiObjectDataOptions struct {
	ShowType                bool `json:"showType"`
	ShowContent             bool `json:"showContent"`
	ShowBcs                 bool `json:"showBcs"`
	ShowOwner               bool `json:"showOwner"`
	ShowPreviousTransaction bool `json:"showPreviousTransaction"`
	ShowStorageRebate       bool `json:"showStorageRebate"`
	ShowDisplay             bool `json:"showDisplay"`
}

type SuiObjectDataFilter map[string]interface{}

type ObjectFilterByPackage struct {
	Package string `json:"Package"`
}

type ObjectFilterByStructType struct {
	StructType string `json:"StructType"`
}

type ObjectFilterByAddressOwner struct {
	AddressOwner string `json:"AddressOwner"`
}

type ObjectFilterByObjectOwner struct {
	ObjectOwner string `json:"ObjectOwner"`
}

type ObjectFilterByObjectId struct {
	ObjectId string `json:"ObjectId"`
}

type ObjectFilterByObjectIds struct {
	ObjectIds []string `json:"ObjectIds"`
}

type ObjectFilterByVersion struct {
	Version string `json:"Version"`
}

type SuiObjectResponseQuery struct {
	// If None, no filter will be applied
	Filter interface{} `json:"filter"`
	// config which fields to include in the response, by default only digest is included
	Options SuiObjectDataOptions `json:"options"`
}

type SuiGetObjectRequest struct {
	// the ID of the queried object
	ObjectId string `json:"ObjectId"`
	// config which fields to include in the response, by default only digest is included
	Options SuiObjectDataOptions `json:"options"`
}

type SuiXGetOwnedObjectsRequest struct {
	// the owner's Sui address
	Address string `json:"address" validate:"checkAddress"`
	// the objects query criteria
	Query SuiObjectResponseQuery
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}

type PaginatedObjectsResponse struct {
	Data        []SuiObjectResponse `json:"data"`
	NextCursor  string              `json:"nextCursor"`
	HasNextPage bool                `json:"hasNextPage"`
}

type SuiObjectResponse struct {
	Data  SuiObjectData          `json:"data"`
	Error SuiObjectResponseError `json:"error"`
}

type SuiObjectResponseError struct {
	Code     string `json:"code"`
	Error    string `json:"error"`
	ObjectId string `json:"object_id"`
	Version  int    `json:"version"`
	Digest   string `json:"digest"`
}

type ObjectOwner struct {
	// the owner's Sui address
	AddressOwner string      `json:"AddressOwner"`
	ObjectOwner  string      `json:"ObjectOwner"`
	Shared       ObjectShare `json:"Shared"`
}

type ObjectShare struct {
	InitialSharedVersion int `json:"initial_shared_version"`
}

type SuiRawMoveObject struct {
	Type              string `json:"type"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Version           int    `json:"version"`
	BcsBytes          string `json:"bcsBytes"`
}

type DisplayFieldsResponse struct {
	Data  interface{}            `json:"data,omitempty"`
	Error SuiObjectResponseError `json:"error,omitempty"`
}

type SuiParsedData struct {
	DataType string `json:"dataType"`
	SuiMoveObject
	SuiMovePackage
}

type SuiObjectData struct {
	ObjectId            string                `json:"objectId"`
	Version             string                `json:"version"`
	Digest              string                `json:"digest"`
	Type                string                `json:"type"`
	Owner               interface{}           `json:"owner"`
	PreviousTransaction string                `json:"previousTransaction"`
	Display             DisplayFieldsResponse `json:"display"`
	Content             SuiParsedData         `json:"content"`
	Bcs                 SuiRawData            `json:"bcs"`
}

type SuiMultiGetObjectsRequest struct {
	ObjectIds []string             `json:"objectIds"`
	Options   SuiObjectDataOptions `json:"options"`
}

type SuiXGetDynamicFieldRequest struct {
	ObjectId string `json:"objectId"`
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
}

type DynamicFieldInfo struct {
	Name       string `json:"name"`
	BcsName    string `json:"bcsName"`
	Type       string `json:"type"`
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Version    int    `json:"version"`
	Digest     string `json:"digest"`
}

type PaginatedDynamicFieldInfoResponse struct {
	Data        []DynamicFieldInfo `json:"data"`
	NextCursor  string             `json:"nextCursor"`
	HasNextPage bool               `json:"hasNextPage"`
}

type SuiXGetDynamicFieldObjectRequest struct {
	ObjectId         string           `json:"objectId"`
	DynamicFieldName DynamicFieldName `json:"dynamicFieldName"`
}

type SuiTryGetPastObjectRequest struct {
	// the ID of the queried object
	ObjectId string `json:"objectId"`
	// the version of the queried object
	Version uint64               `json:"version"`
	Options SuiObjectDataOptions `json:"options"`
}

type PastObjectResponse struct {
	Status  string      `json:"status"`
	Details interface{} `json:"details"`
}

type SuiGetLoadedChildObjectsRequest struct {
	Digest string `json:"digest"`
}

type ChildObjectsResponse struct {
	LoadedChildObjects []*SuiLoadedChildObject `json:"loadedChildObjects"`
}

type SuiLoadedChildObject struct {
	ObjectID       string `json:"objectId"`
	SequenceNumber string `json:"sequenceNumber"`
}

type PastObject struct {
	// the ID of the queried object
	ObjectId string `json:"objectId"`
	// the version of the queried object
	Version string `json:"version"`
}

type SuiTryMultiGetPastObjectsRequest struct {
	// a vector of object and versions to be queried
	MultiGetPastObjects []*PastObject
	// options for specifying the content to be returned
	Options SuiObjectDataOptions
}
