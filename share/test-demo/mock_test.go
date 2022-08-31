package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, nil)

	if v := GetFromDB(m, "Tom"); v != 100 {
		t.Fatal("expected 100, but got", v)
	}
}
