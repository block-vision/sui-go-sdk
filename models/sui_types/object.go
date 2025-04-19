package sui_types

type SuiObjectRef struct {
	Digest   string `json:"digest"`
	ObjectId string `json:"objectId"`
	Version  uint64 `json:"version"`
}

type Owner struct {
	AddressOwner string `json:"addressOwner,omitempty"`
	ObjectOwner  string `json:"objectOwner,omitempty"`
}
