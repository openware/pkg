package nats

import (
	"errors"
	"time"

	"github.com/nats-io/nats.go"
)

type BucketManager interface {
	CreateBucketOrFindExisting(name string, replicas int) (*Bucket, error)
	BucketExists(name string) (bool, error)
	DeleteBucket(name string) error
}

type BucketInterface interface {
	GetValue(key string) ([]byte, error)
	GetKeyCreateDate(key string) (time.Time, error)
	GetAllTheKeys() ([]string, error)
	GetHistoryOfTheKey(key string) ([]KeyVal, error)
	AddPair(key string, val []byte) error
	Delete(key string) error
}

type Manager struct {
	Js nats.JetStreamContext
}

type Bucket struct {
	Items nats.KeyValue
}

type KeyVal struct {
	Key        string
	Value      []byte
	CreateDate time.Time
}

func NewBucketManager(nc *nats.Conn) (*Manager, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &Manager{
		Js: js,
	}, nil
}

func (s *Manager) CreateBucketOrFindExisting(name string, replicas int) (*Bucket, error) {
	keyVal, err := s.Js.KeyValue(name)
	if errors.Is(err, nats.ErrBucketNotFound) {
		keyVal, err = s.Js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket:   name,
			Replicas: replicas,
		})
	} else if err != nil {
		return nil, err
	}

	if keyVal != nil {
		bucket := &Bucket{
			Items: keyVal,
		}

		return bucket, nil
	}

	return nil, err
}

func (s *Manager) BucketExists(name string) (bool, error) {
	_, err := s.Js.KeyValue(name)
	if errors.Is(err, nats.ErrBucketNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Manager) DeleteBucket(name string) error {
	return s.Js.DeleteKeyValue(name)
}

func (b *Bucket) GetValue(key string) ([]byte, error) {
	val, err := b.Items.Get(key)
	if err != nil {
		return nil, err
	}
	return val.Value(), nil
}

func (b *Bucket) GetKeyCreateDate(key string) (time.Time, error) {
	val, err := b.Items.Get(key)
	if err != nil {
		return time.Time{}, err
	}
	return val.Created(), err
}

func (b *Bucket) GetAllTheKeys() ([]string, error) {
	return b.Items.Keys()
}

func (b *Bucket) GetHistoryOfTheKey(key string) ([]KeyVal, error) {
	keyHistory, err := b.Items.History(key)
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

func (b *Bucket) AddPair(key string, val []byte) error {
	_, err := b.Items.Get(key)

	if errors.Is(err, nats.ErrKeyNotFound) {
		_, err = b.Items.Create(key, val)
		return err
	} else if err != nil {
		return err
	}

	_, err = b.Items.Put(key, val)
	return err
}

func (b *Bucket) Delete(key string) error {
	return b.Items.Delete(key)
}
