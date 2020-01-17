//+build windows

package inter

import (
	"fmt"
	"github.com/woshihot/go-lib/utils/os"
	"golang.org/x/sys/windows/registry"
	"log"
)

const (
	ShInterfaceStaticSet = "netsh interface ipv4 set address name=%s source=static addr=%s mask=%s"
	//ShInterfaceDHCPSet   = "netsh interface ipv4 set address name=%s source=dhcp"

	Parameters         = `SYSTEM\CurrentControlSet\services\SharedAccess\Parameters`
	ScopeAddress       = `ScopeAddress`
	ScopeAddressBackup = `ScopeAddressBackup`
)

func interfaceSimple(inter Interface) error {
	ip := ""
	cmd := ""
	if nil == inter.Ip {
		ip = ""
		//cmd = fmt.Sprintf(ShInterfaceDHCPSet, inter.Name)
	} else {
		ip = inter.Ip.String()
		//cmd = fmt.Sprintf(ShInterfaceStaticSet, inter.Name, ip, inter.Mask)
	}
	cmd = fmt.Sprintf(ShInterfaceStaticSet, inter.Name, ip, inter.Mask)
	if len(inter.GateWay) > 0 {
		cmd = cmd + " gateway=" + inter.GateWay + " gwmetric=1"
	} else {
		cmd = cmd + " gateway="
	}

	_, err := os.RunCmd("cmd", "/C", cmd)

	if err != nil {
		log.Printf("[setIp-error] %s\n", err)
		return err
	}

	//修改注册表 scopeAddress
	err = editWinRegistry(registry.LOCAL_MACHINE, Parameters, ScopeAddress, registry.ALL_ACCESS, inter.Ip.String())
	if err != nil {
		log.Printf("[editRegedit-error] %s\n", err)
		return err
	}

	//scopeAddressBackup
	err = editWinRegistry(registry.LOCAL_MACHINE, Parameters, ScopeAddressBackup, registry.ALL_ACCESS, inter.Ip.String())
	if err != nil {
		log.Printf("[editRegedit-error] %s\n", err)
		return err
	}

	//set Dns
	//if len(inter.DNS) > 0 {
	dnsInter := DNSInterface{
		Name: inter.Name,
		DNS:  inter.DNS,
	}

	errDns := setDns(dnsInter)

	if errDns != nil {
		return errDns
	}
	//}

	return nil
}

func interfaceMulti(inters []Interface) error {
	var err error
	for _, inter := range inters {
		err = interfaceSimple(inter)
	}
	return err
}

func editWinRegistry(key registry.Key, path, keyName string, access uint32, val string) error {

	k, err := registry.OpenKey(key, path, access)
	if err != nil {
		log.Printf("[openRegedit-error] %s\n", err)
		return err
	}
	defer k.Close()

	err = k.SetStringValue(keyName, val)
	if err != nil {
		log.Printf("[setRegeditKey-error] %s\n", err)
		return err
	}
	return nil
}
