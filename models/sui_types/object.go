package sui_types

type SuiObjectRef struct {
	Digest   string `json:"digest"`
	ObjectId string `json:"objectId"`
	Version  uint64 `json:"version"`
}

type SuiSharedObject struct {
	ObjectId             string `json:"objectId"`
	InitialSharedVersion uint64 `json:"initialSharedVersion"`
	Mutable              bool   `json:"mutable"`
}

type Owner struct {
	AddressOwner string `json:"addressOwner,omitempty"`
	ObjectOwner  string `json:"objectOwner,omitempty"`
}
