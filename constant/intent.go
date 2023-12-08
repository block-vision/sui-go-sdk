package constant

type IntentScope = uint8

const (
	TransactionDataIntentScope    IntentScope = 0
	TransactionEffectsIntentScope IntentScope = 1
	CheckpointSummaryIntentScope  IntentScope = 2
	PersonalMessageIntentScope    IntentScope = 3
)
