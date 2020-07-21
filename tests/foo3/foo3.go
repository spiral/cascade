package foo3

import (
	"github.com/spiral/cascade/tests/foo2"
	"github.com/spiral/cascade/tests/foo4"
)

type S3 struct {
}

func (s3 *S3) Depends() []interface{} {
	return []interface{}{
		s3.SomeOtherDep,
	}
}

func (s3 *S3) SomeOtherDep(svc *foo4.S4, svc2 foo2.S2) error {
	return nil
}

// Depends on S3
func (s3 *S3) Init(svc foo2.S2) error {
	return nil
}

func (s3 *S3) Serve() chan error {
	errCh := make(chan error, 1)
	return errCh
}

func (s3 *S3) Stop() error {
	return nil
}
