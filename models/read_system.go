package models

type SuiGetCheckpointRequest struct {
	CheckpointID string `json:"id"`
}

type EpochRollingGasCostSummary struct {
	ComputationCost         string `json:"computationCost"`
	StorageCost             string `json:"storageCost"`
	StorageRebate           string `json:"storageRebate"`
	NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
}

type CheckpointResponse struct {
	Epoch                      string                     `json:"epoch"`
	SequenceNumber             string                     `json:"sequenceNumber"`
	Digest                     string                     `json:"digest"`
	NetworkTotalTransactions   string                     `json:"networkTotalTransactions"`
	PreviousDigest             string                     `json:"previousDigest"`
	EpochRollingGasCostSummary EpochRollingGasCostSummary `json:"epochRollingGasCostSummary"`
	TimestampMs                string                     `json:"timestampMs"`
	Transactions               []string                   `json:"transactions"`
	CheckpointCommitments      []interface{}              `json:"checkpointCommitments"`
	ValidatorSignature         string                     `json:"validatorSignature"`
}

type SuiGetCheckpointsRequest struct {
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
	// query result ordering, default to false (ascending order), oldest record first
	DescendingOrder bool `json:"descendingOrder"`
}

type PaginatedCheckpointsResponse struct {
	Data        []CheckpointResponse `json:"data"`
	NextCursor  string               `json:"nextCursor"`
	HasNextPage bool                 `json:"hasNextPage"`
}

type SuiXGetCommitteeInfoRequest struct {
	Epoch string `json:"epoch"`
}

type SuiXGetCommitteeInfoResponse struct {
	Epoch      string     `json:"epoch"`
	Validators [][]string `json:"validators"`
}

type SuiXGetStakesRequest struct {
	Owner string `json:"owner"`
}

type SuiXGetStakesByIdsRequest struct {
	StakedSuiIds []string `json:"stakedSuiIds"`
}

type DelegatedStakeInfo struct {
	StakedSuiId       string `json:"stakedSuiId"`
	StakeRequestEpoch string `json:"stakeRequestEpoch"`
	StakeActiveEpoch  string `json:"stakeActiveEpoch"`
	Principal         string `json:"principal"`
	Status            string `json:"status"`
	EstimatedReward   string `json:"estimatedReward"`
}

type DelegatedStakesResponse struct {
	ValidatorAddress string               `json:"validatorAddress"`
	StakingPool      string               `json:"stakingPool"`
	Stakes           []DelegatedStakeInfo `json:"stakes"`
}

type SuiXGetEpochsRequest struct {
	// optional paging cursor
	Cursor interface{} `json:"cursor"`
	// maximum number of items per page
	Limit uint64 `json:"limit" validate:"lte=50"`
	// query result ordering, default to false (ascending order), oldest record first
	DescendingOrder bool `json:"descendingOrder"`
}

type SuiValidatorSummary struct {
	SuiAddress                   SuiAddress `json:"suiAddress"`
	ProtocolPubkeyBytes          string     `json:"protocolPubkeyBytes"`
	NetworkPubkeyBytes           string     `json:"networkPubkeyBytes"`
	WorkerPubkeyBytes            string     `json:"workerPubkeyBytes"`
	ProofOfPossessionBytes       string     `json:"proofOfPossessionBytes"`
	OperationCapId               string     `json:"operationCapId"`
	Name                         string     `json:"name"`
	Description                  string     `json:"description"`
	ImageUrl                     string     `json:"imageUrl"`
	ProjectUrl                   string     `json:"projectUrl"`
	P2pAddress                   string     `json:"p2pAddress"`
	NetAddress                   string     `json:"netAddress"`
	PrimaryAddress               string     `json:"primaryAddress"`
	WorkerAddress                string     `json:"workerAddress"`
	NextEpochProtocolPubkeyBytes string     `json:"nextEpochProtocolPubkeyBytes"`
	NextEpochProofOfPossession   string     `json:"nextEpochProofOfPossession"`
	NextEpochNetworkPubkeyBytes  string     `json:"nextEpochNetworkPubkeyBytes"`
	NextEpochWorkerPubkeyBytes   string     `json:"nextEpochWorkerPubkeyBytes"`
	NextEpochNetAddress          string     `json:"nextEpochNetAddress"`
	NextEpochP2pAddress          string     `json:"nextEpochP2pAddress"`
	NextEpochPrimaryAddress      string     `json:"nextEpochPrimaryAddress"`
	NextEpochWorkerAddress       string     `json:"nextEpochWorkerAddress"`
	VotingPower                  string     `json:"votingPower"`
	GasPrice                     string     `json:"gasPrice"`
	CommissionRate               string     `json:"commissionRate"`
	NextEpochStake               string     `json:"nextEpochStake"`
	NextEpochGasPrice            string     `json:"nextEpochGasPrice"`
	NextEpochCommissionRate      string     `json:"nextEpochCommissionRate"`
	StakingPoolId                string     `json:"stakingPoolId"`
	StakingPoolActivationEpoch   string     `json:"stakingPoolActivationEpoch"`
	StakingPoolDeactivationEpoch string     `json:"stakingPoolDeactivationEpoch"`
	StakingPoolSuiBalance        string     `json:"stakingPoolSuiBalance"`
	RewardsPool                  string     `json:"rewardsPool"`
	PoolTokenBalance             string     `json:"poolTokenBalance"`
	PendingStake                 string     `json:"pendingStake"`
	PendingPoolTokenWithdraw     string     `json:"pendingPoolTokenWithdraw"`
	PendingTotalSuiWithdraw      string     `json:"pendingTotalSuiWithdraw"`
	ExchangeRatesId              string     `json:"exchangeRatesId"`
	ExchangeRatesSize            string     `json:"exchangeRatesSize"`
}

type EndOfEpochInfo struct {
	LastCheckpointId             string `json:"lastCheckpointId"`
	EpochEndTimestamp            string `json:"epochEndTimestamp"`
	ProtocolVersion              string `json:"protocolVersion"`
	ReferenceGasPrice            string `json:"referenceGasPrice"`
	TotalStake                   string `json:"totalStake"`
	StorageFundReinvestment      string `json:"storageFundReinvestment"`
	StorageCharge                string `json:"storageCharge"`
	StorageRebate                string `json:"storageRebate"`
	StorageFundBalance           string `json:"storageFundBalance"`
	StakeSubsidyAmount           string `json:"stakeSubsidyAmount"`
	TotalGasFees                 string `json:"totalGasFees"`
	TotalStakeRewardsDistributed string `json:"totalStakeRewardsDistributed"`
	LeftoverStorageFundInflow    string `json:"leftoverStorageFundInflow"`
}

type EpochInfo struct {
	Epoch                  string                `json:"epoch"`
	Validators             []SuiValidatorSummary `json:"validators"`
	EpochTotalTransactions string                `json:"epochTotalTransactions"`
	FirstCheckpointId      string                `json:"firstCheckpointId"`
	EpochStartTimestamp    string                `json:"epochStartTimestamp"`
	EndOfEpochInfo         EndOfEpochInfo        `json:"endOfEpochInfo"`
}

type PaginatedEpochInfoResponse struct {
	Data        []EpochInfo `json:"data"`
	NextCursor  string      `json:"nextCursor"`
	HasNextPage bool        `json:"hasNextPage"`
}

type SuiSystemStateSummary struct {
	Epoch                                 string                `json:"epoch"`
	ProtocolVersion                       string                `json:"protocolVersion"`
	SystemStateVersion                    string                `json:"systemStateVersion"`
	StorageFundTotalObjectStorageRebates  string                `json:"storageFundTotalObjectStorageRebates"`
	StorageFundNonRefundableBalance       string                `json:"storageFundNonRefundableBalance"`
	ReferenceGasPrice                     string                `json:"referenceGasPrice"`
	SafeMode                              bool                  `json:"safeMode"`
	SafeModeStorageRewards                string                `json:"safeModeStorageRewards"`
	SafeModeComputationRewards            string                `json:"safeModeComputationRewards"`
	SafeModeStorageRebates                string                `json:"safeModeStorageRebates"`
	SafeModeNonRefundableStorageFee       string                `json:"safeModeNonRefundableStorageFee"`
	EpochStartTimestampMs                 string                `json:"epochStartTimestampMs"`
	EpochDurationMs                       string                `json:"epochDurationMs"`
	StakeSubsidyStartEpoch                string                `json:"stakeSubsidyStartEpoch"`
	MaxValidatorCount                     string                `json:"maxValidatorCount"`
	MinValidatorJoiningStake              string                `json:"minValidatorJoiningStake"`
	ValidatorLowStakeThreshold            string                `json:"validatorLowStakeThreshold"`
	ValidatorVeryLowStakeThreshold        string                `json:"validatorVeryLowStakeThreshold"`
	ValidatorLowStakeGracePeriod          string                `json:"validatorLowStakeGracePeriod"`
	StakeSubsidyBalance                   string                `json:"stakeSubsidyBalance"`
	StakeSubsidyDistributionCounter       string                `json:"stakeSubsidyDistributionCounter"`
	StakeSubsidyCurrentDistributionAmount string                `json:"stakeSubsidyCurrentDistributionAmount"`
	StakeSubsidyPeriodLength              string                `json:"stakeSubsidyPeriodLength"`
	StakeSubsidyDecreaseRate              int                   `json:"stakeSubsidyDecreaseRate"`
	TotalStake                            string                `json:"totalStake"`
	ActiveValidators                      []SuiValidatorSummary `json:"activeValidators"`
	PendingActiveValidatorsId             string                `json:"pendingActiveValidatorsId"`
	PendingActiveValidatorsSize           string                `json:"pendingActiveValidatorsSize"`
	PendingRemovals                       []string              `json:"pendingRemovals"`
	StakingPoolMappingsId                 string                `json:"stakingPoolMappingsId"`
	StakingPoolMappingsSize               string                `json:"stakingPoolMappingsSize"`
	InactivePoolsId                       string                `json:"inactivePoolsId"`
	InactivePoolsSize                     string                `json:"inactivePoolsSize"`
	ValidatorCandidatesId                 string                `json:"validatorCandidatesId"`
	ValidatorCandidatesSize               string                `json:"validatorCandidatesSize"`
	AtRiskValidators                      []string              `json:"atRiskValidators"`
	ValidatorReportRecords                [][]interface{}       `json:"validatorReportRecords"`
}

type ValidatorsApy struct {
	Apys  []Apy  `json:"apys"`
	Epoch string `json:"epoch"`
}

type Apy struct {
	Address string  `json:"address"`
	Apy     float64 `json:"apy"`
}

type SuiGetProtocolConfigRequest struct {
	Version string `json:"version"`
}

type ProtocolConfigResponse struct {
	MinSupportedProtocolVersion string                       `json:"minSupportedProtocolVersion"`
	MaxSupportedProtocolVersion string                       `json:"maxSupportedProtocolVersion"`
	ProtocolVersion             string                       `json:"protocolVersion"`
	FeatureFlags                map[string]bool              `json:"featureFlags"`
	Attributes                  map[string]map[string]string `json:"attributes"`
}
