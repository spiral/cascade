package ServeRetryErr

import (
	"errors"
	"time"
)

type S1ServeErr struct {
}

// No deps
func (s *S1ServeErr) Init(s2 *S2) error {
	return nil
}

func (s *S1ServeErr) Serve() chan error {
	errCh := make(chan error, 1)
	go func() {
		time.Sleep(time.Millisecond * 300)
		errCh <- errors.New("test serve error")
	}()
	return errCh
}

func (s *S1ServeErr) Stop() error {
	return nil
}