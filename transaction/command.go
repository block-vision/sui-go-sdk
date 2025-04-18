package transaction

func moveCall(input ProgrammableMoveCall) Command {
	return MoveCall{
		Value: input,
	}
}

func transferObjects(input TransferObjectsValue) Command {
	return TransferObjects{
		Value: input,
	}
}

func splitCoins(input SplitCoinsValue) Command {
	return SplitCoins{
		Value: input,
	}
}

func mergeCoins(input MergeCoinsValue) Command {
	return MergeCoins{
		Value: input,
	}
}

func publish(input PublishValue) Command {
	return Publish{
		Value: input,
	}
}

func makeMoveVec(input MakeMoveVecValue) Command {
	return MakeMoveVec{
		Value: input,
	}
}

func upgrade(input UpgradeValue) Command {
	return Upgrade{
		Value: input,
	}
}
