package myfile

import (
	"os"
	"path/filepath"
)

func FileExits(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func DirectoryExists(path string) bool {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func CreateOrAppendFile(filePath string) (err error) {
	if !FileExits(filePath) {
		if !DirectoryExists(filePath) {
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}
