package repository

import (
	"errors"
	"testing"
	"time"
)

func TestAuditBatchFailureCannotBlockSender(t *testing.T) {
	errorsOut := make(chan error)
	start := make(chan struct{})
	returned := make(chan struct{}, 2)
	for i := 0; i < 2; i++ {
		go func() {
			<-start
			DeliverAuditError(errors.New("database unavailable"), errorsOut)
			returned <- struct{}{}
		}()
	}
	close(start)
	for i := 0; i < 2; i++ {
		select {
		case <-returned:
		case <-time.After(100 * time.Millisecond):
			t.Fatal("audit error sender blocked without a receiver")
		}
	}
}
