package mock

import "bytes"

type MockReadWriteCloser struct {
	B bytes.Buffer
}

func (m *MockReadWriteCloser) Read(p []byte) (n int, err error) {
	return m.B.Read(p)
}

func (m *MockReadWriteCloser) Write(p []byte) (n int, err error) {
	return m.B.Write(p)
}

func (m *MockReadWriteCloser) Close() error {
	return nil
}
