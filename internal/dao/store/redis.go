package store

import "io"

type RedisStore struct{}

func (r *RedisStore) Store(name string, data io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) Get(name string) (io.Reader, error) {
	//TODO implement me
	panic("implement me")
}
