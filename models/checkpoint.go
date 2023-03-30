package models

type GetLatestCheckpointSequenceNumberRequest struct {
}

type GetCheckpointSummaryRequest struct {
	SequenceNumber uint64 `json:"sequenceNumber,omitempty"`
}

type GetCheckpointSummaryResponse struct {
	Epoch                      uint64      `json:"epoch,omitempty"`
	SequenceNumber             uint64      `json:"sequence_number,omitempty"`
	NetworkTotalTransaction    uint64      `json:"network_total_transactions,omitempty"`
	ContentDigest              string      `json:"content_digest,omitempty"`
	PreviousDigest             string      `json:"previous_digest,omitempty"`
	EpochRollingGasCostSummary interface{} `json:"epoch_rolling_gas_cost_summary,omitempty"`
}

type GetCheckpointContentsRequest struct {
	Digest string `json:"digest,omitempty"`
}

type GetCheckpointContentsResponse struct {
	Transactions []interface{} `json:"transactions,omitempty"`
}

type GetCheckpointContentsBySequenceNumberRequest struct {
	SequenceNumber uint64
}

type GetCheckpointContentsBySequenceNumberResponse struct {
	Transactions []interface{} `json:"transactions,omitempty"`
}
