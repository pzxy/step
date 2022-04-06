package main

import (
	"os"
	"path"
)

func main() {
	// 创建目录和文件
	initDbFile("./asd/")
}

func initDbFile(dbPath string) error {
	var err error
	if err = os.MkdirAll(path.Dir(dbPath), os.ModePerm); err != nil {
		return err
	}
	if _, err = os.Stat(dbPath); err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	if _, err = os.Create(dbPath); err != nil {
		return err
	}
	return nil
}
