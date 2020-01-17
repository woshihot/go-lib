//+build windows

package inter

import (
	"errors"
	"fmt"
	"github.com/woshihot/go-lib/utils/os"
	"strings"
)

func GetMacByIp(screen string) (string, error) {
	if "" == screen {
		return "", errors.New("ip can not empty")
	}
	arpResult, err := os.RunCmd("cmd", "/C", fmt.Sprintf("arp -a %s", screen))
	if err != nil {
		return "", err
	}
	macs := findMac(arpResult)
	if nil != macs && len(macs) > 0 {
		return strings.Replace(macs[0], "-", ":", -1), nil
	}
	return "", nil
}
