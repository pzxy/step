package tmp

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"step/grammar/codeskill/log"
	"testing"
)

func TestGlogWrap(t *testing.T) {
	log.InitGLog()

	if err := sum(); err != nil {
		err = errors.WithStack(err)
		glog.Errorf("%+v", err)
	}
}

func sum() error {
	return errors.New("Wrong param!")
}
