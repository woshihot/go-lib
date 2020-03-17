package file

import (
	"os"
)

func WriteByte(path string, content []byte) error {
	_, err := OpenFile(path)
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
