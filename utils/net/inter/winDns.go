//+build windows

package inter

import (
	"fmt"
	"github.com/woshihot/go-lib/utils/os"
	"golang.org/x/sys/windows/registry"
	"log"
	"regexp"
)

const (
	ShGetDns              = "netsh interface ipv4 sh dnsservers name=%s"
	ShSetDns              = "netsh interface ipv4 set dns name=%s source=static addr=%s register=primary"
	StandaloneDhcpAddress = `StandaloneDhcpAddress`
)

func getAllDns(names []string) []DNSInterface {

	var result []DNSInterface
	if nil == names || len(names) == 0 {
		return nil
	}
	for _, name := range names {
		dns, err := os.RunCmd("cmd", "/C", fmt.Sprintf(ShGetDns, name))
		if nil != err {
			continue
		}
		result = append(result, DNSInterface{name, findDNS(dns)})
	}

	return result
}

func getDns(name string) DNSInterface {
	var result DNSInterface

	dns, err := os.RunCmd("cmd", "/C", fmt.Sprintf(ShGetDns, name))
	if nil != err {
		return result
	}
	return DNSInterface{name, findDNS(dns)}
}

func findDNS(input string) []string {
	regEx := regexp.MustCompile(RegexpIp4)
	return regEx.FindAllString(input, -1)
}

func setDns(dns DNSInterface) error {

	/*
	  修改单个DNS 如果dns地址多个的话，会默认修改为第一个地址
	*/

	var dns0 string

	if len(dns.DNS[0]) > 0 {
		dns0 = dns.DNS[0]
	} else {
		dns0 = " " + "\"\""
	}

	_, err := os.RunCmd("cmd", "/C", fmt.Sprintf(ShSetDns, dns.Name, dns0))

	if err != nil {
		log.Printf("[setLanIp-error] %s\n", err)
		return err
	}

	//standAloneDhcpAddress
	err = editWinRegistry(registry.LOCAL_MACHINE, Parameters, StandaloneDhcpAddress, registry.ALL_ACCESS, dns.DNS[0])
	if err != nil {
		log.Printf("[editRegedit-error] %s\n", err)
		return err
	}

	return nil
}
