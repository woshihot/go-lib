//+build linux

package inter

import (
	"github.com/woshihot/go-lib/utils/file"
	"strings"
)

const (
	InterfaceFilePath = "/etc/network/interfaces"
	InterfaceFlag     = "auto"
	NameServerFlag    = "dns-nameservers"
)

func getAllDns(names []string) []DNSInterface {

	lines := file.ReadFileLine(InterfaceFilePath)
	var result []DNSInterface
	if nil == names || len(names) == 0 {
		return nil
	}
	var name string
	var dns string
	for _, line := range lines {
		if strings.HasPrefix(line, InterfaceFlag) {
			towardsName := strings.Replace(strings.TrimSpace(strings.TrimPrefix(line, InterfaceFlag)), "\n", "", -1)
			if Contains(names, towardsName) {
				name = towardsName
			}
		}
		if strings.HasPrefix(line, NameServerFlag) && "" != name {
			dns = strings.Replace(strings.TrimSpace(strings.TrimPrefix(line, NameServerFlag)), "\n", "", -1)
			result = append(result, DNSInterface{name, strings.Split(dns, " ")})
		}
	}

	return result
}

func getDns(name string) DNSInterface {
	var result DNSInterface
	dnss := getAllDns([]string{name})
	if nil == dnss || len(dnss) < 1 {
		return result
	}
	result = DNSInterface{name, dnss[0].DNS}
	return result
}
