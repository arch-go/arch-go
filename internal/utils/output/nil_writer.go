package output

type nilWriter struct{}

func CreateNilWriter() nilWriter {
	return nilWriter{}
}

func (n nilWriter) Write(_ []byte) (int, error) {
	return 0, nil
}
