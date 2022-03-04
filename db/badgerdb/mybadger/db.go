package mybadger

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

type BadgerDB struct {
	db *badger.DB
}

func NewClient(dir string) (c *BadgerDB, err error) {
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		log.Fatal(err)
	}
	bg := &BadgerDB{
		db: db,
	}
	return bg, nil
}

func (b *BadgerDB) PutString(k, v string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(k), []byte(v))
	})
}
