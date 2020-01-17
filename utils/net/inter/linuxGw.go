//+build linux

package inter

import (
	"fmt"
	"github.com/woshihot/go-lib/utils/os"
	"strings"
)

const (
	SH_GET_GW = "/sbin/route -n|grep UG |grep %s |awk '{print $2}'| uniq"
)

func getGwByInter(name string) (string, error) {
	result, err := os.RunCmd("/bin/bash", "-c", fmt.Sprintf(SH_GET_GW, name))
	if nil != err {
		return "", err
	} else {
		return strings.TrimSpace(result), nil
	}
}
