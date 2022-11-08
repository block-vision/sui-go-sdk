package sui_error

import "errors"

var (
	ErrInvalidJson            = errors.New("invalid json response")
	ErrUnknownSignatureScheme = errors.New("unknown scheme sign scheme flag")
	ErrInvalidEncryptFlag     = errors.New("invalid encrypt flag")
	ErrNoKeyStoreInfo         = errors.New("no keystore info, make sure already loaded sui.keystore")
	ErrAddressNotInKeyStore   = errors.New("address not in keystore, make sure already loaded sui.keystore")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidWebsocketClient = errors.New("invalid websocket client")
)
