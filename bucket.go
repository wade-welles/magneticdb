package magneticdb

import (
	"bytes"
	"errors"
)

var (
	errBucketIsNotExist = errors.New("Bucket is not exist")
	errKeyIsNotFound    = errors.New("Key is not found")
	errBucketExist      = errors.New("Bucket already exist")
)

// BucketConfig provides configuration for each bucket.
// Optional parameter for CreateBucket
type BucketConfig struct {
	keysize   uint
	valuesize uint
}

type Bucket struct {
	items     map[string][]*Item
	keysize   uint
	valuesize uint
}

// New provides creational of the new bucket
func NewBucket() *Bucket {
	b := new(Bucket)
	b.items = map[string][]*Item{}
	return b
}

// CreateBucket provides creational of the new bucket
func (b *Bucket) CreateBucket(title string, cfg *BucketConfig) error {
	if cfg != nil {
		if cfg.keysize != 0 {
			b.keysize = cfg.keysize
		}

		if cfg.valuesize != 0 {
			b.valuesize = cfg.valuesize
		}
	}
	_, ok := b.items[title]
	if ok {
		return errBucketExist
	}
	b.items[title] = []*Item{}
	return nil
}

func (b *Bucket) SetToBucket(title string, key, value []byte) error {
	_, ok := b.items[title]
	if !ok {
		return errBucketIsNotExist
	}

	if b.keysize != 0 && b.keysize > uint(len(key)) {

	}

	newitem, err := set(title, key, value)
	if err != nil {
		return err
	}
	b.items[title] = append(b.items[title], newitem)
	return nil
}

func (b *Bucket) GetFromBucket(title string, key []byte) ([]byte, error) {
	items, ok := b.items[title]
	if !ok {
		return nil, errBucketIsNotExist
	}

	for _, item := range items {
		if bytes.Equal(key, item.key) {
			return item.value, nil
		}
	}

	return nil, errKeyIsNotFound
}
