package main

import (
	"context"
	"github.com/networm6/gopherBox/route"
	"github.com/networm6/gopherBox/tunnel"
)

func main() {
	var a uint64 = 0
	var b uint64 = 0
	tun := tunnel.NewTunnel(context.Background(), &a, &b)
	tun.SetConf(&tunnel.TunConfig{
		DeviceName: "SimonFuck",
		MTU:        1500,
		BufferSize: 64 * 1024,
		MixinFunc: func(bytes []byte) []byte {
			return bytes
		},
		CIDRv4: "172.16.0.25/24",
		CIDRv6: "",
	})
	r := route.Route{
		Server:         false,
		LocalGatewayV4: "",
		LocalGatewayV6: "",
		CIDRv4:         "",
		CIDRv6:         "",
		ServerTunIPv4:  "",
		ServerTunIPv6:  "",
	}
	tun.SetCallBack(func(conf tunnel.TunConfig) {
		_ = r.SetRoute(conf.DeviceName, "", conf.MTU)
	}, func() {
		_ = r.ResetRoute("")
	})
	tun.Start()
	tun.Destroy()
}
