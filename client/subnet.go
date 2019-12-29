package client

import (
	"net"
)

// Subnet represents a subnet with its local IP and C.I.D.R.
type Subnet struct {
	name  string
	ip    net.IP
	net   *net.IPNet
	iface *Interface
}

// NewSubnet returns an instance of a Subnet.
// The local IP and route are applied for the interface.
func NewSubnet(iface *Interface, name, ip, network string) (*Subnet, error) {
	s := Subnet{name: name, iface: iface}

	s.ip = net.ParseIP(ip)
	_, s.net, _ = net.ParseCIDR(network)

	if err := iface.AddIP(s.ip); err != nil {
		return nil, err
	}

	if err := iface.AddRoute(s.net); err != nil {
		return nil, err
	}

	return &s, nil
}

// SubnetList is a simple list of all subnets.
type SubnetList struct {
	subnets []*Subnet
}

// Add adds a subnet to the list.
func (sl *SubnetList) Add(s *Subnet) {
	sl.subnets = append(sl.subnets, s)
}

// LocalIPs returns a list of the local IPs for all subnets.
func (sl *SubnetList) LocalIPs() []net.IP {
	ips := []net.IP{}
	for _, s := range sl.subnets {
		ips = append(ips, s.ip)
	}
	return ips
}

// LocalNets returns a list of the local networks for all subnets.
func (sl *SubnetList) LocalNets() []*net.IPNet {
	nets := []*net.IPNet{}
	for _, s := range sl.subnets {
		nets = append(nets, s.net)
	}
	return nets
}
