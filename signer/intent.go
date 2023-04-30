package signer

type AppId int

const (
	Sui AppId = 0
)

type IntentVersion int

const (
	V0 IntentVersion = 0
)

type IntentScope int

const (
	TransactionData    IntentScope = 0
	TransactionEffects IntentScope = 1
	CheckpointSummary  IntentScope = 2
	PersonalMessage    IntentScope = 3
)

func IntentWithScope(intentScope IntentScope) []int {
	return []int{int(intentScope), int(V0), int(Sui)}
}

