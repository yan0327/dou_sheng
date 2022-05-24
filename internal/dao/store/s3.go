package store

import (
	"github.com/minio/minio-go"
	"io"
	"sync"
)

// S3Store 兼容Amazon S3接口的对象存储
type S3Store struct {
	endpoint        string
	accessKeyId     string
	secretAccessKey string
	bucketName      string
	pool            *sync.Pool
}

func MakeS3Store(endpoint string, accessKeyId string, secretAccessKey string, bucketName string) *S3Store {
	return &S3Store{
		endpoint,
		accessKeyId,
		secretAccessKey,
		bucketName,
		nil,
	}
}

func MakeS3PoolStore(endpoint string, accessKeyId string, secretAccessKey string, bucketName string) *S3Store {
	s := &S3Store{
		endpoint,
		accessKeyId,
		secretAccessKey,
		bucketName,
		&sync.Pool{
			New: func() interface{} {
				client, _ := minio.New(endpoint, accessKeyId, secretAccessKey, false)
				return client
			},
		},
	}
	for i := 0; i < 20; i++ {
		client, _ := minio.New(endpoint, accessKeyId, secretAccessKey, false)
		s.pool.Put(client)
	}
	return s
}

func (m *S3Store) getClient() (*minio.Client, error) {
	if m.pool != nil {
		return m.pool.Get().(*minio.Client), nil
	}
	client, err := minio.New(m.endpoint, m.accessKeyId, m.secretAccessKey, false)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (m *S3Store) putBackClient(c *minio.Client) {
	if m.pool != nil {
		m.pool.Put(c)
	}
}

func (m *S3Store) Store(name string, data io.Reader) error {
	client, err := m.getClient()
	defer m.putBackClient(client)
	if err != nil {
		return err
	}
	_, err = client.PutObject(m.bucketName, name, data, -1, minio.PutObjectOptions{})
	return err
}

func (m *S3Store) Delete(name string) error {
	client, err := m.getClient()
	defer m.putBackClient(client)
	if err != nil {
		return err
	}
	return client.RemoveObject(m.bucketName, name)
}

func (m *S3Store) Get(name string) (io.Reader, error) {
	client, err := m.getClient()
	defer m.putBackClient(client)
	if err != nil {
		return nil, err
	}
	obj, err := client.GetObject(m.bucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}
