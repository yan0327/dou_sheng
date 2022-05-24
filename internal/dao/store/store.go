package store

import "io"

type Store interface {
	Store(name string, data io.Reader) error
	Delete(name string) error
	Get(name string) (io.Reader, error)
}
