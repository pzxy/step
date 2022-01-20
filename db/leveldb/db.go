package leveldb

import (
	"context"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"time"
)

type LevelDB struct {
	db *leveldb.DB
}

func NewClient(dir string) (c *LevelDB, errEdgeX error) {
	client, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		if client, err = leveldb.RecoverFile(dir, nil); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	ldb := &LevelDB{
		db: client,
	}
	return ldb, nil
}

func (l *LevelDB) PutString(key, value string, wo *opt.WriteOptions) error {
	return l.db.Put([]byte(key), []byte(value), wo)
}

func (l *LevelDB) GetString(key string, wo *opt.ReadOptions) (string, error) {
	b, err := l.db.Get([]byte(key), wo)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func (l *LevelDB) PutStringExpire(key, value string, dur time.Duration) error {

	return l.db.Put([]byte(key), []byte(value), nil)
}

func (l *LevelDB) GetStringExpire(key string, dur time.Duration) (string, error) {
	b, err := l.db.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func (l *LevelDB) Demo() {
	l.PutString("1", "a", nil)
}

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func formatMs(ctx context.Context, dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		internal.Logger.Printf(
			ctx,
			"specified duration is %s, but minimal supported value is %s - truncating to 1ms",
			dur, time.Millisecond,
		)
		return 1
	}
	return int64(dur / time.Millisecond)
}

func formatSec(ctx context.Context, dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		internal.Logger.Printf(
			ctx,
			"specified duration is %s, but minimal supported value is %s - truncating to 1s",
			dur, time.Second,
		)
		return 1
	}
	return int64(dur / time.Second)
}
