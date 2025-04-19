package transaction

import "errors"

var (
	ErrSignerNotSet           = errors.New("signer not set")
	ErrSenderNotSet           = errors.New("sender not set")
	ErrGasDataNotFullySet     = errors.New("gas data not fully set")
	ErrObjectIdNotSet         = errors.New("object id not set")
	ErrObjectTypeNotSupported = errors.New("object type not supported")
	ErrInvalidSuiAddress      = errors.New("invalid Sui address")
)
