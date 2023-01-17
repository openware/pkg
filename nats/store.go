package nats

import (
	"errors"
	"time"

	"github.com/nats-io/nats.go"
)

type BucketManager interface {
	CreateBucketOrFindExisting(name string) (*bucket, error)
	BucketExists(name string) (bool, error)
	DeleteBucket(name string) error
}

type Bucket interface {
	GetValue(key string) ([]byte, error)
	GetKeyCreateDate(key string) (time.Time, error)
	GetAllTheKeys() ([]string, error)
	GetHistoryOfTheKey(key string) ([]KeyVal, error)
	AddPair(key string, val []byte) error
	Delete(key string) error
}

type bucketManager struct {
	js nats.JetStreamContext
}

type bucket struct {
	items nats.KeyValue
}

type KeyVal struct {
	Key        string
	Value      []byte
	CreateDate time.Time
}

func NewBucketManager(nc *nats.Conn) (*bucketManager, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &bucketManager{
		js: js,
	}, nil
}

func (s *bucketManager) CreateBucketOrFindExisting(name string, replicas int) (*bucket, error) {
	keyVal, err := s.js.KeyValue(name)
	if errors.Is(err, nats.ErrBucketNotFound) {
		keyVal, err = s.js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket:   name,
			Replicas: replicas,
		})
	} else if err != nil {
		return nil, err
	}

	if keyVal != nil {
		bucket := &bucket{
			items: keyVal,
		}

		return bucket, nil
	}

	return nil, err
}

func (s *bucketManager) BucketExists(name string) (bool, error) {
	_, err := s.js.KeyValue(name)
	if errors.Is(err, nats.ErrBucketNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *bucketManager) DeleteBucket(name string) error {
	return s.js.DeleteKeyValue(name)
}

func (b *bucket) GetValue(key string) ([]byte, error) {
	val, err := b.items.Get(key)
	if err != nil {
		return nil, err
	}
	return val.Value(), nil
}

func (b *bucket) GetKeyCreateDate(key string) (time.Time, error) {
	val, err := b.items.Get(key)
	if err != nil {
		return time.Time{}, err
	}
	return val.Created(), err
}

func (b *bucket) GetAllTheKeys() ([]string, error) {
	return b.items.Keys()
}

func (b *bucket) GetHistoryOfTheKey(key string) ([]KeyVal, error) {
	keyHistory, err := b.items.History(key)
	if err != nil {
		return nil, err
	}

	history := make([]KeyVal, len(keyHistory))
	for i := 0; i < len(keyHistory); i++ {
		history[i] = KeyVal{
			Key:        keyHistory[i].Key(),
			Value:      keyHistory[i].Value(),
			CreateDate: keyHistory[i].Created(),
		}
	}

	return history, nil
}

func (b *bucket) AddPair(key string, val []byte) error {
	_, err := b.items.Get(key)

	if errors.Is(err, nats.ErrKeyNotFound) {
		_, err = b.items.Create(key, val)
		return err
	} else if err != nil {
		return err
	}

	_, err = b.items.Put(key, val)
	return err
}

func (b *bucket) Delete(key string) error {
	return b.items.Delete(key)
}
