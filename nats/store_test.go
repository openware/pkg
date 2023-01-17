package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func initBucketManager() *bucketManager {
	//nc, _ := InitNats("localhost:4222")
	nc, _ := InitEmbededNats()
	bm, _ := NewBucketManager(nc)

	return bm
}

func TestCreateBucketOrFindExisting(t *testing.T) {
	bm := initBucketManager()

	exist, err := bm.BucketExists("foo")
	assert.NoError(t, err)
	assert.Equal(t, false, exist)

	bucket, err := bm.CreateBucketOrFindExisting("foo", 1)
	assert.NoError(t, err)
	assert.NotNil(t, bucket)

	exist, err = bm.BucketExists("foo")
	assert.NoError(t, err)
	assert.Equal(t, true, exist)
}

func TestDeleteBucket(t *testing.T) {
	bm := initBucketManager()
	bucket, _ := bm.CreateBucketOrFindExisting("foo", 1)
	assert.NotNil(t, bucket)

	bm.DeleteBucket("foo")
	exist, _ := bm.BucketExists("foo")
	assert.Equal(t, false, exist)
}

func TestGetSetValue(t *testing.T) {
	bm := initBucketManager()
	bucket, _ := bm.CreateBucketOrFindExisting("foo", 1)

	data, err := bucket.GetValue("baz")
	assert.ErrorIs(t, err, nats.ErrKeyNotFound)
	assert.Nil(t, data)

	err = bucket.AddPair("baz", []byte("bar"))
	assert.NoError(t, err)

	data, _ = bucket.GetValue("baz")
	assert.Equal(t, []byte("bar"), data)
}

func TestGetAllTheKeys(t *testing.T) {

}
