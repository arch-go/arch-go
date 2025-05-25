package pkgone

import (
	jsonAlias "encoding/json"
	"io"

	"github.com/arch-go/arch-go/internal/verifications/naming/test/pkgtwo"
	"github.com/stretchr/testify/require"
)

type Service interface {
	NormalFunc(s []string) (int, error)
	EmbeddedGenerics[int]
	pkgtwo.OtherPackageGenerics[int]
	EmbeddedNormal
	EmbeddedOtherPackageGenerics
	io.Writer
	jsonAlias.Marshaler
}

type EmbeddedNormal interface {
	EmbeddedNormalFunc()
}

type EmbeddedGenerics[T any] interface {
	EmbeddedGenericsFunc(t T) (T, error)
}

type EmbeddedOtherPackageGenerics interface {
	pkgtwo.OtherPackage
}

var _ Service = new(DefaultService)

type DefaultService struct{}

func (d DefaultService) OtherPackageEmbeddedNormalFunc() {}

func (d DefaultService) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (d DefaultService) EmbeddedGenericsFunc(_ int) (int, error) {
	return 0, nil
}

func (d DefaultService) OtherPackageEmbeddedGenericsFunc(_ int) {}

func (d DefaultService) EmbeddedNormalFunc() {}

func (d DefaultService) NormalFunc(_ []string) (int, error) {
	return 0, nil
}

func (d DefaultService) Write(_ []byte) (int, error) {
	return 0, nil
}

type SomeInterface interface {
	SomeMethod()
}

type SomeType struct{}

func (st SomeType) SomeMethod() {}

type ServiceWithGenerics[T SomeInterface] interface {
	NormalFunc(s []string) (int, error)
	EmbeddedGenerics[T]
	pkgtwo.OtherPackageGenerics[T]
	EmbeddedNormal
	EmbeddedOtherPackageGenerics
	io.Writer
	jsonAlias.Marshaler
}

var _ ServiceWithGenerics[SomeType] = new(DefaultServiceWithGenerics[SomeType])

type DefaultServiceWithGenerics[T SomeInterface] struct{}

func (d *DefaultServiceWithGenerics[T]) NormalFunc(_ []string) (int, error) {
	return 0, nil
}

func (d *DefaultServiceWithGenerics[T]) EmbeddedGenericsFunc(_ T) (T, error) {
	return *new(T), nil
}

func (d *DefaultServiceWithGenerics[T]) OtherPackageEmbeddedGenericsFunc(_ T) {}

func (d *DefaultServiceWithGenerics[T]) EmbeddedNormalFunc() {}

func (d *DefaultServiceWithGenerics[T]) OtherPackageEmbeddedNormalFunc() {}

func (d *DefaultServiceWithGenerics[T]) Write(_ []byte) (int, error) {
	return 0, nil
}

func (d *DefaultServiceWithGenerics[T]) MarshalJSON() ([]byte, error) {
	return nil, nil
}

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
