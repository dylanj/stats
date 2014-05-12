package stats

import (
	"bytes"
	"io"
	"os"
)

type FileOpener interface {
	Open(filename string) (io.ReadCloser, error)
	Create(filename string) (io.WriteCloser, error)
}

type nilFileOpener struct{}
type osFileOpener struct{}
type fakeFileOpener struct {
	*bytes.Buffer
}

// Open
func (_ *nilFileOpener) Open(name string) (io.ReadCloser, error) {
	return &os.File{}, os.ErrNotExist
}

// Create
func (_ *nilFileOpener) Create(name string) (io.WriteCloser, error) {
	return nil, nil
}

// Open
func (o osFileOpener) Open(name string) (io.ReadCloser, error) {
	return os.Open(name)
}

// Create
func (o osFileOpener) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

// Open
func (o *fakeFileOpener) Open(name string) (io.ReadCloser, error) {
	return o, nil
}

// Create
func (o *fakeFileOpener) Create(name string) (io.WriteCloser, error) {
	return o, nil
}

// Close
func (o *fakeFileOpener) Close() error {
	return nil
}
