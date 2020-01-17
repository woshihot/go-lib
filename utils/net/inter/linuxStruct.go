package inter

import (
	"fmt"
	"net"
	"strings"
)

func (i *Interface) String() string {
	if i.IsEmpty() {
		return ""
	}
	return splicing(i.Head(), i.Foot()) + "\n"
}

func (i *Interface) Head() string {
	if i.IsPhysical || i.IsLoopback() {

		var name, iface string
		name = fmt.Sprintf("auto %s", i.Name)
		iface = fmt.Sprintf("iface %s inet %s", i.Name, getTypeByIp(i.Ip /*, isDhcp(i)*/))
		return splicing(name, iface)
	}
	return ""
}
func (i *Interface) Foot() string {
	if !i.IsPhysical || i.IsLoopback() {
		return ""
	}
	var address, mask, gateWay, dns string

	//if !isDhcp(i) {
	if nil != i.Ip {
		address = fmt.Sprintf("address %s", i.Ip.String())
	}
	if "" != i.Mask {
		mask = fmt.Sprintf("netmask %s", i.Mask)
	}
	//}
	if "" != i.GateWay {
		gateWay = fmt.Sprintf("gateway %s", i.GateWay)
	}

	if nil != i.DNS && len(i.DNS) > 0 {
		dns = "dns-nameservers"
		suffix := ""
		for _, d := range i.DNS {
			suffix = suffix + " " + d
		}
		if "" == strings.TrimSpace(suffix) {
			dns = ""
		} else {
			dns = dns + suffix
		}
	}
	return splicing(address, mask, gateWay, dns)

}

func getTypeByIp(ip net.IP /*, isDhcp bool*/) string {
	/*if nil == ip || isDhcp {
		return "dhcp"
	} else*/if ip.IsLoopback() {
		return "loopback"
	} else {
		return "static"
	}
}
func splicing(args ...string) string {
	var result = ""
	for i, arg := range args {
		if "" != arg {
			result += arg
			if i != len(args)-1 {
				result += "\n"
			}
		}
	}
	return result
}

//func isDhcp(i *Interface) bool {
//	return nil == i.Ip || isDhcpNow(i.Name)
//}

//func isDhcpNow(name string) bool {
//	cmd := "ps aux | grep dhcp |grep -v grep|grep %s"
//	result, err := os.RunCommand(fmt.Sprintf(cmd, name))
//	if nil != err {
//		return false
//	} else {
//		return "" != result
//	}
//}
