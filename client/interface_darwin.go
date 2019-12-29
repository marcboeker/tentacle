package client

import (
	"net"
	"os/exec"

	"github.com/songgao/water"
)

// SetupInterface creates a new TUN interface on macOS hosts.
func SetupInterface() (*Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}
	iface, err := water.New(config)
	if err != nil {
		return nil, err
	}

	return &Interface{iface}, nil
}

// AddIP adds an IP address to the new interface.
func (i Interface) AddIP(ip net.IP) error {
	cmd := exec.Command("ifconfig", i.Interface.Name(), ip.String(), ip.String(), "up")
	return cmd.Run()
}

// AddRoute creates a new route to the new interface.
func (i Interface) AddRoute(cidr *net.IPNet) error {
	cmd := exec.Command("route", "-n", "add", "-net", cidr.String(), "-interface", i.Interface.Name())
	return cmd.Run()
}
