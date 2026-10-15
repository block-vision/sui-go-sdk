package transaction

import "errors"

var (
	ErrSignerNotSet         = errors.New("signer not set")
	ErrSenderNotSet         = errors.New("sender not set")
	ErrSuiClientNotSet      = errors.New("sui client not set")
	ErrGasDataNotAllSet     = errors.New("gas data not all set")
	ErrGasPriceNotSet       = errors.New("gas price not set: call SetGasPrice or attach a SuiClient")
	ErrInvalidSuiAddress    = errors.New("invalid sui address")
	ErrInvalidObjectId      = errors.New("invalid object id")
	ErrObjectNotSupportType = errors.New("object not support type")

	ErrUnsupportedTransactionKind        = errors.New("unsupported transaction kind: only programmable transactions can be encoded or decoded")
	ErrUnresolvedInput                   = errors.New("unresolved input: pass a full object reference (CallArg with ObjectArg) instead of an object id")
	ErrInvalidTypeTag                    = errors.New("invalid type tag")
	ErrAddressBalanceGasNeedsValidDuring = errors.New("empty gas payment pays gas from an address balance and requires a ValidDuring or Validity expiration")
)
