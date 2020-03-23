//Package config properties util
package config

import (
	"github.com/magiconair/properties"
	"github.com/woshihot/go-lib/utils/file"
	"github.com/woshihot/go-lib/utils/log"
	"sync"
)

var TAG_ERROR = "[CONFIG_ERROR]"

type config struct {
	*properties.Properties
	lock     *sync.RWMutex
	wlock    *sync.RWMutex
	FilePath []string
}

//NewConfig create a config pointer
func NewConfig() *config {
	return &config{
		lock:  new(sync.RWMutex),
		wlock: new(sync.RWMutex)}
}

//LoadPath load file or url
func (c *config) LoadPath(paths ...string) error {
	var err error

	for _, p := range paths {
		file.OpenFile(p)
	}

	c.Properties, err = properties.LoadFiles(paths, properties.UTF8, true)
	if err != nil {
		return err
	}
	c.FilePath = paths
	properties.ErrorHandler = func(e error) {
		log.EF(TAG_ERROR, "properties errorHandler path = %s, %v\n", paths, e)
	}
	return err
}
