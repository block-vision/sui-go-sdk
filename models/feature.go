package models

type MoveCallAndExecuteTransactionRequest struct {
	MoveCallRequest
}

type MoveCallAndExecuteTransactionResponse struct {
	ExecuteTransactionResponse
}

type BatchAndExecuteTransactionRequest struct {
	BatchTransactionRequest
}

type BatchAndExecuteTransactionResponse struct {
	Result []ExecuteTransactionResponse
}
