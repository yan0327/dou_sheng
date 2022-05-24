package store

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var s3store Store
var s3poolStore Store

func TestMain(m *testing.M) {
	s3store = MakeS3Store(
		"localhost:9000",
		"minioadmin",
		"minioadmin",
		"douyin",
	)
	s3poolStore = MakeS3PoolStore(
		"localhost:9000",
		"minioadmin",
		"minioadmin",
		"douyin",
	)
	os.Exit(m.Run())
}

func TestS3Store(t *testing.T) {
	err := s3store.Store("1", bytes.NewReader([]byte("abc")))
	assert.Nil(t, err)
	obj, err := s3store.Get("1")
	data, _ := ioutil.ReadAll(obj)
	assert.Equal(t, data, []byte("abc"))
}

func BenchCase(s Store) {
	_ = s.Store("1", bytes.NewReader([]byte("abc")))
	obj, _ := s.Get("1")
	ioutil.ReadAll(obj)
}

func BenchmarkS3Store(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchCase(s3store)
	}
}

func BenchmarkS3PoolStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchCase(s3poolStore)
	}
}

func BenchmarkS3StoreParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			BenchCase(s3store)
		}
	})
}

func BenchmarkS3PoolStoreParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			BenchCase(s3poolStore)
		}
	})
}
