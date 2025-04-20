package transaction

import "errors"

var (
	ErrSignerNotSet       = errors.New("signer not set")
	ErrSenderNotSet       = errors.New("sender not set")
	ErrSuiClientNotSet    = errors.New("sui client not set")
	ErrGasDataNotFullySet = errors.New("gas data not fully set")
	ErrInvalidSuiAddress  = errors.New("invalid sui address")
	ErrInvalidObjectId    = errors.New("invalid object id")
)
