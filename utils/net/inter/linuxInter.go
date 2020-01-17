//+build linux

package inter

import (
	"github.com/woshihot/go-lib/utils/file"
)

const (
	StartContent = "source-directory /etc/network/interfaces.d"

	//EndContent     = "pre-up iptables-restore < /etc/iptables/rules.v4"
	InterfacesPath = "/etc/network/interfaces"
)

func interfaceSimple(inter Interface) error {
	interAll, err := GetInterfaces()
	if err != nil {
		return err
	}
	var interfaces string
	interfaces = StartContent + "\n"
	var isAdd = true
	for _, interSimple := range interAll {
		if interSimple.Name == inter.Name {
			isAdd = false
			interfaces = interfaces + inter.String() + "\n"
		} else if "" != interSimple.String() {
			interfaces = interfaces + interSimple.String() + "\n"
		}
	}
	if isAdd {
		interfaces = interfaces + inter.String() + "\n"
	}
	//if os.UBUNTU_16 == os.GetUbuntuVersion() {
	//	interfaces = interfaces + EndContent + "\n"
	//}
	return rewriteInterface(interfaces)
}

func rewriteInterface(interfaces string) error {
	//log.Printf(interfaces)
	//return nil
	return file.WriteByte(InterfacesPath, []byte(interfaces))
}

func interfaceMulti(inters []Interface) error {
	interAll, err := GetInterfaces()
	if err != nil {
		return err
	}
	var interfaces string
	interfaces = StartContent + "\n"
	for _, interSimple := range interAll {
		_, result := isChangeInterface(inters, interSimple)
		if "" != result.String() {
			interfaces = interfaces + result.String() + "\n"
		}
	}

	return rewriteInterface(interfaces)
}

func isChangeInterface(inters []Interface, inter Interface) (bool, Interface) {
	var isChange bool
	var result Interface
	for _, interSimple := range inters {
		if interSimple.Name == inter.Name {
			isChange = true
			result = interSimple
		}
	}
	if !isChange {
		result = inter
	}
	return isChange, result
}
