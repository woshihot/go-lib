package inter

import "net"

func GetLocalMac(interName string) (string, error) {
	inter, err := net.InterfaceByName(interName)
	if err != nil {
		return "", err
	}
	return inter.HardwareAddr.String(), nil
}
