package transaction

import "github.com/samber/lo"

func moveCall(input ProgrammableMoveCall) Command {
	return Command{
		MoveCall: lo.ToPtr(input),
	}
}

func transferObjects(input TransferObjects) Command {
	return Command{
		TransferObjects: lo.ToPtr(input),
	}
}

func splitCoins(input SplitCoins) Command {
	return Command{
		SplitCoins: lo.ToPtr(input),
	}
}

func mergeCoins(input MergeCoins) Command {
	return Command{
		MergeCoins: lo.ToPtr(input),
	}
}

func publish(input Publish) Command {
	return Command{
		Publish: lo.ToPtr(input),
	}
}

func makeMoveVec(input MakeMoveVec) Command {
	return Command{
		MakeMoveVec: lo.ToPtr(input),
	}
}

func upgrade(input Upgrade) Command {
	return Command{
		Upgrade: lo.ToPtr(input),
	}
}
