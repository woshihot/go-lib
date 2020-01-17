package inter

import (
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
)

const (
	RegexpBlock   = "(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)"
	RegexpIp4     = RegexpBlock + "\\." + RegexpBlock + "\\." + RegexpBlock + "\\." + RegexpBlock
	RegexpMask    = "\\d*"
	RegexpIp4Mask = RegexpIp4 + "/" + RegexpMask
	RegexpMac     = "(([a-f0-9]{2}:)|([a-f0-9]{2}-)){5}[a-f0-9]{2}"
)

func GetInterfaces(names ...string) ([]Interface, error) {

	var inters []net.Interface
	var err error
	var errInters []string
	if nil == names || len(names) == 0 {
		inters, err = net.Interfaces()
		if err != nil {
			log.Printf("GetInterfaces err = %s\n", err.Error())
			return nil, err
		}
		names = interfaceNames(inters)
	} else {
		for _, name := range names {
			inter, errInter := net.InterfaceByName(name)
			if errInter != nil {
				errInters = append(errInters, name)
				continue
			}
			inters = append(inters, *inter)
		}
	}
	if len(inters) != len(names) {
		err = errors.New(fmt.Sprintf("%v is not find", errInters))
	}
	allDns := getAllDns(names)
	return netInterToLocals(inters, allDns), err
}

func GetInterface(name string) (Interface, error) {
	var result Interface
	inter, err := net.InterfaceByName(name)
	if nil != err {
		return result, err
	}
	dns := getDns(name)
	return netInterToLocal(inter, dns.DNS), nil
}

func SetInterfaces(inters ...Interface) error {
	if len(inters) == 1 {
		return setInterfaceSample(inters[0])
	} else if len(inters) > 0 {
		return setInterfaceMulti(inters)
	}
	return errors.New("nil interfaces")
}

func setInterfaceSample(inter Interface) error {
	return interfaceSimple(inter)
}

func setInterfaceMulti(inters []Interface) error {
	return interfaceMulti(inters)
}

func netInterToLocals(inters []net.Interface, allDns []DNSInterface) []Interface {
	var result []Interface
	for _, inter := range inters {

		dns := findDns(inter.Name, allDns)
		added := netInterToLocal(&inter, dns)
		result = append(result, added)
	}
	return result
}
func netInterToLocal(inter *net.Interface, dns []string) Interface {
	var result Interface
	ip, mask, err := getIpMask(inter)
	if err != nil || ip == nil || "" == mask {
		return result
	}
	gateWay, err := getGwByInter(inter.Name)
	if err != nil {
		gateWay = ""
	}
	mac := inter.HardwareAddr.String()
	isPhysical := "" != mac
	return Interface{
		Name:       inter.Name,
		Flag:       inter.Flags,
		Ip:         ip,
		Mask:       mask,
		GateWay:    gateWay,
		DNS:        dns,
		Mac:        mac,
		IsPhysical: isPhysical,
	}
}

func maskToString(mask net.IPMask) string {
	var result = ""
	for _, b := range mask {
		if "" != result {
			result = result + "."
		}
		result = fmt.Sprint(result, b)
	}
	return result
}

func interfaceNames(inters []net.Interface) []string {
	var result []string
	for _, inter := range inters {
		result = append(result, inter.Name)
	}
	return result
}

func getIpMask(inter *net.Interface) (net.IP, string, error) {
	address, err := inter.Addrs()
	if err != nil {
		return nil, "", err
	}
	var ip net.IP
	var mask string
	for _, addr := range address {
		ipWithMask := addr.String()
		isIp4, err := regexp.MatchString(RegexpIp4Mask, ipWithMask)
		if err != nil {
			continue
		}
		if isIp4 {

			ip4, ip4Net, err := net.ParseCIDR(ipWithMask)
			if err != nil {
				continue
			}
			ip = ip4
			mask = maskToString(ip4Net.Mask)
		}
	}
	return ip, mask, nil
}

func findDns(name string, allDns []DNSInterface) []string {
	var dns []string
	if nil != allDns {
		for _, d := range allDns {
			if d.Name == name {
				dns = d.DNS
			}
		}
	}
	if len(dns) > 0 {
		return dns
	} else {
		return nil
	}
}

func Contains(exprs []string, e string) bool {
	for _, expr := range exprs {
		if e == expr {
			return true
		}
	}
	return false
}

func findMac(input string) []string {
	regEx := regexp.MustCompile(RegexpMac)
	return regEx.FindAllString(input, -1)
}
