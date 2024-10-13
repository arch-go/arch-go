package pkgtwo

type OtherPackageGenerics[T any] interface {
	OtherPackageEmbeddedGenericsFunc(t T)
}

type OtherPackage interface {
	OtherPackageEmbeddedNormalFunc()
}
