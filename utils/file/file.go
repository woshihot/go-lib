package file

import (
	"github.com/woshihot/go-lib/utils/log"
	"os"
	"path/filepath"
)

var TAG_ERROR = "[File-error]"
var TAG_DEBUG = "[File]"

func OpenFile(path string) (*os.File, error) {
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
	} else {
		return os.Open(path)
	}
	return result, nil
}
