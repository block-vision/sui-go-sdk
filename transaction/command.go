package transaction

func moveCall(input ProgrammableMoveCall) Command {
	return Command{
		MoveCall: &input,
	}
}

func transferObjects(input TransferObjects) Command {
	return Command{
		TransferObjects: &input,
	}
}

func splitCoins(input SplitCoins) Command {
	return Command{
		SplitCoins: &input,
	}
}

func mergeCoins(input MergeCoins) Command {
	return Command{
		MergeCoins: &input,
	}
}

func publish(input Publish) Command {
	return Command{
		Publish: &input,
	}
}

func makeMoveVec(input MakeMoveVec) Command {
	return Command{
		MakeMoveVec: &input,
	}
}

func upgrade(input Upgrade) Command {
	return Command{
		Upgrade: &input,
	}
}
