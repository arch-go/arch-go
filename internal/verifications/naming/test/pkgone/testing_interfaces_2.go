package pkgone

import (
	"io"

	"github.com/stretchr/testify/require"
)

type SomeErr struct{}

var _ error = new(SomeErr)

func (e *SomeErr) Error() string {
	return ""
}

type SomeReader struct{}

var _ io.Reader = new(SomeReader)

func (s SomeReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

type ImplementExternalInterface struct{}

var _ require.TestingT = new(ImplementExternalInterface)

func (t *ImplementExternalInterface) Errorf(format string, args ...interface{}) {}

func (t *ImplementExternalInterface) FailNow() {}
