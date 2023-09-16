package queue

import "context"

type Queue interface {
	Receive(ctx context.Context) ([]string, error)
	Send(ctx context.Context, msg string) error
}
