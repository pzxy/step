package main

// go get -u github.com/golang/mock/gomock
// go get -u github.com/golang/mock/mockgen

//go:generate mockgen -source=mock.go -destination=gen_mock.go -package=main

type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
