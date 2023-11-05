//go:build windows
// +build windows

package route

import (
	"errors"
	"github.com/networm6/gopherBox/netbox"
	"github.com/networm6/gopherBox/osbox"
)

func (route *Route) SetRoute(interfaceName, serverAddr string, MTU int) error {
	ServerTunIPv4 := route.ServerTunIPv4
	ServerTunIPv6 := route.ServerTunIPv6
	LocalGatewayV4 := route.LocalGatewayV4
	LocalGatewayV6 := route.LocalGatewayV6
	commandPath := "cmd"

	if route.Server {
		return nil
	}
	serverAddrIP := netbox.LookupServerAddrIP(serverAddr)
	if serverAddrIP == nil {
		return errors.New("setRoute fail. Cannot access server")
	}

	if LocalGatewayV4 != "" {
		osbox.ExecCmd(commandPath, "/C", "route", "delete", "0.0.0.0", "mask", "0.0.0.0")
		osbox.ExecCmd(commandPath, "/C", "route", "add", "0.0.0.0", "mask", "0.0.0.0", ServerTunIPv4, "metric", "6")
		if serverAddrIP.To4() != nil {
			osbox.ExecCmd(commandPath, "/C", "route", "add", serverAddrIP.To4().String()+"/32", LocalGatewayV4, "metric", "5")
		}
	}
	if LocalGatewayV6 != "" {
		osbox.ExecCmd(commandPath, "/C", "route", "-6", "delete", "::/0", "mask", "::/0")
		osbox.ExecCmd(commandPath, "/C", "route", "-6", "add", "::/0", "mask", "::/0", ServerTunIPv6, "metric", "6")
		if serverAddrIP.To16() != nil {
			osbox.ExecCmd(commandPath, "/C", "route", "-6", "add", serverAddrIP.To16().String()+"/128", LocalGatewayV6, "metric", "5")
		}
	}
	return nil
}

func (route *Route) ResetRoute(serverAddr string) error {
	LocalGatewayV4 := route.LocalGatewayV4
	LocalGatewayV6 := route.LocalGatewayV6
	commandPath := "cmd"
	if route.Server {
		return nil
	}
	serverAddrIP := netbox.LookupServerAddrIP(serverAddr)
	if serverAddrIP == nil {
		return errors.New("setRoute fail. Cannot access server")
	}

	if LocalGatewayV4 != "" {
		osbox.ExecCmd(commandPath, "/C", "route", "delete", "0.0.0.0", "mask", "0.0.0.0")
		osbox.ExecCmd(commandPath, "/C", "route", "add", "0.0.0.0", "mask", "0.0.0.0", LocalGatewayV4, "metric", "6")
	}
	if LocalGatewayV6 != "" {
		osbox.ExecCmd(commandPath, "/C", "route", "-6", "delete", "::/0", "mask", "::/0")
		osbox.ExecCmd(commandPath, "/C", "route", "-6", "add", "::/0", "mask", "::/0", LocalGatewayV6, "metric", "6")
	}
	return nil
}
