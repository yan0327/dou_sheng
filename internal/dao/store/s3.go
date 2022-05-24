package store

import "io"

// S3Store 兼容Amazon S3接口的对象存储
type S3Store struct{}

func (m *S3Store) Store(name string, data io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (m *S3Store) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}

func (m *S3Store) Get(name string) (io.Reader, error) {
	//TODO implement me
	panic("implement me")
}
