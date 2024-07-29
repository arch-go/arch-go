package output

import "io"

type nilWriter struct{}

func CreateNilWriter() io.Writer {
	return nilWriter{}
}

func (n nilWriter) Write(_ []byte) (int, error) {
	return 0, nil
}
