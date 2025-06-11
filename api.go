package goroutine

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrPanicRecover = errors.New("panic recover")
)

func Run(ctx context.Context, f func(ctx context.Context)) {
	SafeGo(func() { f(ctx) })
}

func SafeGo(runner func()) {
	defer Recovery()
	runner()
}

func RunWithErrChan(ctx context.Context, f func(ctx context.Context) error, errChan chan error) {
	go func(ctx context.Context) {
		defer RecoveryWithErrChan(errChan)
		err := f(ctx)
		errChan <- err
	}(ctx)
}
