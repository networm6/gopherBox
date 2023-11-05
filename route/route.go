package route

type routeInterface interface {
	SetRoute(string, string, int) error
	ResetRoute(string)
}

type Route struct {
	routeInterface
	Server bool

	LocalGatewayV4 string
	LocalGatewayV6 string
	CIDRv4         string
	CIDRv6         string
	ServerTunIPv4  string
	ServerTunIPv6  string
}
