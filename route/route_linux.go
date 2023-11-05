//go:build linux
// +build linux

package route

import (
	"errors"
	"github.com/networm6/gopherBox/netbox"
	"github.com/networm6/gopherBox/osbox"
)

func (route *Route) SetRoute(interfaceName, serverAddr string, MTU int) error {
	CIDRv4 := route.CIDRv4
	CIDRv6 := route.CIDRv6
	LocalGatewayV4 := route.LocalGatewayV4
	LocalGatewayV6 := route.LocalGatewayV6
	commandPath := "/sbin/ip"

	osbox.ExecCmd(commandPath, "link", "set", "dev", interfaceName, "mtu", strconv.Itoa(MTU))
	osbox.ExecCmd(commandPath, "addr", "add", CIDRv4, "dev", interfaceName)
	osbox.ExecCmd(commandPath, "-6", "addr", "add", CIDRv6, "dev", interfaceName)
	osbox.ExecCmd(commandPath, "link", "set", "dev", interfaceName, "up")

	if route.Server {
		return nil
	}

	physicalInterface := netbox.GetInterface()
	serverAddrIP := netbox.LookupServerAddrIP(serverAddr)
	if physicalInterface == "" {
		return errors.New("setRoute fail. Not found physical interface")
	}
	if serverAddrIP == nil {
		return errors.New("setRoute fail. Cannot access server")
	}

	if LocalGatewayV4 != "" {
		osbox.ExecCmd(commandPath, "route", "add", "0.0.0.0/1", "dev", interfaceName)
		osbox.ExecCmd(commandPath, "route", "add", "128.0.0.0/1", "dev", interfaceName)
		if serverAddrIP.To4() != nil {
			osbox.ExecCmd(commandPath, "route", "add", serverAddrIP.To4().String()+"/32", "via", LocalGatewayV4, "dev", physicalInterface)
		}
	}
	if LocalGatewayV6 != "" {
		osbox.ExecCmd(commandPath, "-6", "route", "add", "::/1", "dev", interfaceName)
		if serverAddrIP.To16() != nil {
			osbox.ExecCmd(commandPath, "-6", "route", "add", serverAddrIP.To16().String()+"/128", "via", LocalGatewayV6, "dev", physicalInterface)
		}
	}
	return nil
}

func (route *Route) ResetRoute(serverAddr string) {
	LocalGatewayV4 := route.LocalGatewayV4
	LocalGatewayV6 := route.LocalGatewayV6

	if route.Server {
		return
	}
}
