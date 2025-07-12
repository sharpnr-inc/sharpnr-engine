package interfaces

import "context"

type BaseOperation[T any] interface {
	Create(ctx context.Context, entity *T) error
}
