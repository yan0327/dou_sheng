package store

import (
	"bufio"
	"io"
	"os"
	"path"
)

// FileStore 本地文件存储
type FileStore struct {
	saveDir string
}

func MakeFileStore(saveDir string) *FileStore {
	return &FileStore{saveDir: saveDir}
}

func (f *FileStore) Store(name string, data io.Reader) error {
	dir := path.Join(f.saveDir, name)
	out, err := os.Create(dir)
	defer out.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, data)
	return err
}

func (f *FileStore) Delete(name string) error {
	dir := path.Join(f.saveDir, name)
	return os.Remove(dir)
}

func (f *FileStore) Get(name string) (io.Reader, error) {
	dir := path.Join(f.saveDir, name)
	file, err := os.Open(dir)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}
