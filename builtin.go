package goroutine

import (
	"runtime"

	"code.byted.org/gopkg/logs"
	"github.com/pkg/errors"
)

func Recovery() {
	if r := recover(); r != nil {
		buf := make([]byte, 1<<18)
		buf = buf[:runtime.Stack(buf, false)]
		logs.Errorf("recover, stack=%v", buf)
	}
}

func RecoveryWithErrChan(errChan chan error) {
	if r := recover(); r != nil {
		buf := make([]byte, 1<<18)
		buf = buf[:runtime.Stack(buf, false)]
		logs.Errorf("recover, stack=%v", buf)
		errChan <- errors.Wrap(ErrPanicRecover, string(buf))
	}
}
