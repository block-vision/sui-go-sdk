package transaction

import "errors"

const (
	defaultGasBudget = 50000000
)

var (
	ErrSignerNotSet           = errors.New("signer not set")
	ErrObjectIdNotSet         = errors.New("object id not set")
	ErrObjectTypeNotSupported = errors.New("object type not supported")
)
