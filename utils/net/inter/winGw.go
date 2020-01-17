//+build windows

package inter

import (
	"bytes"
	"errors"
	"net"
	"os"
	"syscall"
	"unsafe"
)

func getGwByInter(name string) (string, error) {

	interfaces, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}

	aList, err := getWinAdapterList()
	if err != nil {
		return "", err
	}
	for ai := aList; ai != nil; ai = ai.Next {

		if interfaces.Index == int(ai.Index) {
			gwl := &ai.GatewayList
			gw := string(bytes.Trim(gwl.IpAddress.String[:], "\x00"))
			return gw, nil
		}
	}
	return "", errors.New("not find gateway")
}

func getWinAdapterList() (*syscall.IpAdapterInfo, error) {

	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}
