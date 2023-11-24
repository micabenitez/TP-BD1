package lib

import (
	fmt "fmt"

	bolt "go.etcd.io/bbolt"
)

func WriteToBucket(bucketName string, key []byte, val []byte) error {
	db := ConnectionBolt()
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	err = b.Put(key, val)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ReadUniqueFromBucket(bucketName string, key []byte) ([]byte, error) {
	db := ConnectionBolt()
	defer db.Close()
	var buf []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf, err
}

func SelectAllFromBucket(bucketName string) {
	db := ConnectionBolt()
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucketName))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil

		})

		return nil
	})

}
