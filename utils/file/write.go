package file

import (
	"github.com/woshihot/go-lib/utils/log"
	"os"
	"path/filepath"
)

func WriteByte(path string, content []byte) error {
	_, err := CreateFile(path)
	if nil != err {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := f.Write(content); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func CreateFile(path string) (*os.File, error) {
	var result *os.File
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		dir := filepath.Dir(path)
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
		if result, err = os.Create(path); err != nil {
			log.EF(TAG_ERROR, "CreateFile -> %s\n", err.Error())
			return nil, err
		}
	}
	return result, nil
}
