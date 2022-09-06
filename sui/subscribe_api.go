package sui

import "context"

type ISubscribeAPI interface {
	// SubscribeEvent TODO
	SubscribeEvent(ctx context.Context, filter interface{}) (<-chan interface{}, error)
}
