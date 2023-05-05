package database

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/etcd-io/bbolt"
)

// Push: This function adds an item to the top of the stack.

// Pop: This function removes and returns the item at the top of the stack.

// Size: This function returns the number of items in the stack.

// DB...
var db *bbolt.DB

// Queue
type Queue struct {
	db *bolt.DB
}

var Que *Queue

// InitDatabase creates a new stack
func InitDatabase(path string) error {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("queue"))
		return err
	})
	if err != nil {
		return err
	}
	Que = &Queue{db}
	return nil
}

// Push ...
func (q *Queue) Push(payload []byte) error {
	err := q.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("queue"))
		id, _ := b.NextSequence()
		return b.Put(itob(id), payload)
	})
	if err != nil {
		return err
	}
	return nil
}

// Pop
func (q *Queue) Pop() ([]byte, error) {
	var payload []byte
	err := q.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("queue"))
		c := b.Cursor()
		k, v := c.First()
		if v == nil {
			return fmt.Errorf("queue is empty")
		}
		payload = make([]byte, len(v))
		copy(payload, v)
		return b.Delete(k)
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (q *Queue) Size() int {
	var size int
	q.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("queue"))
		size = b.Stats().KeyN
		return nil
	})
	return size
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}
