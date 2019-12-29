package client

import (
	"fmt"
	"net"

	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
)

// SetupInterface creates a new TUN interface on Linux hosts.
func SetupInterface() (*Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}
	iface, err := water.New(config)
	if err != nil {
		return nil, err
	}

	link, _ := netlink.LinkByName(iface.Name())
	if err := netlink.LinkSetUp(link); err != nil {
		return nil, err
	}

	return &Interface{iface}, nil
}

// AddIP adds an IP address to the new interface.
func (i Interface) AddIP(ip net.IP) error {
	link, _ := netlink.LinkByName(i.Interface.Name())
	addr, _ := netlink.ParseAddr(fmt.Sprintf("%s/32", ip))
	return netlink.AddrAdd(link, addr)
}

// AddRoute creates a new route to the new interface.
func (i Interface) AddRoute(cidr *net.IPNet) error {
	link, _ := netlink.LinkByName(i.Interface.Name())
	route := netlink.Route{LinkIndex: link.Attrs().Index, Dst: cidr}
	return netlink.RouteAdd(&route)
}
