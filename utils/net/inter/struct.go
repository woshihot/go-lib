package inter

import (
	"net"
)

type Interface struct {
	//interface 名称
	Name string

	//interface flag
	Flag net.Flags

	//interface ipv4
	Ip net.IP

	//interface  ipv4子网掩码
	Mask string

	//interface 网关
	GateWay string

	//interface DNS
	DNS []string

	//是否是物理网卡
	IsPhysical bool

	//interface mac only has data in isPhysical = true
	Mac string
}

type DNSInterface struct {
	Name string
	DNS  []string
}

func (i *Interface) IsLoopback() bool {
	if nil != i.Ip {
		return i.Ip.IsLoopback()
	} else {
		return false
	}
}

func (i *Interface) IsEmpty() bool {
	if nil == i.Ip || "" == i.Name {
		return true
	}
	return false
}
