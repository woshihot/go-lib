package str

import (
	"github.com/woshihot/go-lib/log"
	"time"
)

func TimeFormat(t time.Time) string {
	r := t.Format("2006-01-02 15:04:05.999")
	log.Df("TimeFormat time = %s\n", r)
	return r
}
