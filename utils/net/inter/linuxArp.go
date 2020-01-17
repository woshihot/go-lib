//+build linux

package inter

import (
	"errors"
	"fmt"
	"isesol.com/agentServer/util/os"
	"strings"
)

func GetMacByIp(screen string) (string, error) {
	if "" == screen {
		return "", errors.New("ip can not empty")
	}
	mac, err := os.RunCmd("/bin/bash", "-c", fmt.Sprintf("arp -n |grep %s|awk '{print $3}'", screen))
	if err != nil {
		return "", err
	}
	mac = strings.Replace(strings.TrimSpace(mac), "\n", "", -1)

	macs := findMac(mac)
	if nil != macs && len(macs) > 0 {
		return macs[0], nil
	}
	return "", nil
}
