package models

type GetDelegatedStakesRequest struct {
	Owner string `json:"owner,omitempty"`
}

type GetDelegatedStakeResponse struct {
	DelegateStake []interface{}
}

type GetValidatorsRequest struct {
}

type GetValidatorsResponse struct {
	Result []ValidatorMetadata `json:"result"`
}

type ValidatorMetadata struct {
	SuiAddress              string `json:"sui_address"`
	PubkeyBytes             []byte `json:"pubkey_bytes"`
	NetworkPubkeyBytes      []byte `json:"network_pubkey_bytes"`
	WorkerPubkeyBytes       []byte `json:"worker_pubkey_bytes"`
	ProofOfPossessionBytes  []byte `json:"proof_of_possession_bytes"`
	Name                    []byte `json:"name"`
	NetAddress              []byte `json:"net_address"`
	ConsensusAddress        []byte `json:"consensus_address"`
	WorkerAddress           []byte `json:"worker_address"`
	NextEpochStake          uint64 `json:"next_epoch_stake"`
	NextEpochDelegation     uint64 `json:"next_epoch_delegation"`
	NextEpochGasPrice       uint64 `json:"next_epoch_gas_price"`
	NextEpochCommissionRate uint64 `json:"next_epoch_commission_rate"`
}

type GetCommitteeInfoRequest struct {
	EpochId uint64 `json:"epoch"`
}

type GetCommitteeInfoResponse struct {
	Epoch         uint64      `json:"epoch"`
	CommittedInfo interface{} `json:"committee_info"`
}

type GetSuiSystemStateRequest struct {
}

type GetSuiSystemStateResponse struct {
	Info                   interface{} `json:"info"`
	ChainID                interface{} `json:"chain_id"`
	Epoch                  uint64      `json:"epoch"`
	Validators             interface{} `json:"validators"`
	TreasuryCap            interface{} `json:"treasury_cap"`
	StorageFund            interface{} `json:"storage_fund"`
	Parameters             interface{} `json:"parameters"`
	ReferenceGasPrice      uint64      `json:"reference_gas_price"`
	ValidatorReportRecords interface{} `json:"validator_report_records"`
	StakeSubsidy           interface{} `json:"stake_subsidy"`
}

type GetCheckpointRequest struct {
	Id string `json:"id"`
}

type GetCheckpointResponse struct {
	SequenceNumber           string   `json:"sequenceNumber"`
	Digest                   string   `json:"digest"`
	NetworkTotalTransactions uint64   `json:"networkTotalTransactions"`
	TimestampMs              uint64   `json:"timestampMs"`
	PreviousDigest           string   `json:"previousDigest"`
	Transactions             []string `json:"transactions"`
}
