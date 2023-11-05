package tunnel

import (
	"context"
	"fmt"
	"github.com/net-byte/water"
	"github.com/networm6/gopherBox/lifecycle"
)

// Tunnel 结构体
type Tunnel struct {
	lifecycle.LifeInterface
	_ctx    *context.Context
	_cancel *context.CancelFunc

	_conf         *TunConfig
	_tunInterface *water.Interface

	_startCB   func(config TunConfig)
	_destroyCB func()

	_totalReadBytes    *uint64
	_totalWrittenBytes *uint64

	InputStream  chan []byte
	OutputStream chan []byte
}

// NewTunnel 创建。
func NewTunnel(parentCtx context.Context, readBytes, writtenBytes *uint64) *Tunnel {
	ctx, cancel := context.WithCancel(parentCtx)
	tunnel := &Tunnel{
		_ctx:               &ctx,
		_cancel:            &cancel,
		_totalReadBytes:    readBytes,
		_totalWrittenBytes: writtenBytes,
		InputStream:        make(chan []byte),
		OutputStream:       make(chan []byte),
	}
	return tunnel
}

func (tun *Tunnel) SetCallBack(startCB func(config TunConfig), destroyCB func()) {
	tun._startCB = startCB
	tun._destroyCB = destroyCB
}
func (tun *Tunnel) SetConf(conf *TunConfig) error {
	tun._conf = conf
	return tun.createTunnelInterface()
}

func (tun *Tunnel) Start() {
	tun._startCB(*tun._conf)
	go tun.readFromTunnel()
	go tun.writeToTunnel()
}

func (tun *Tunnel) Destroy() {
	(*tun._cancel)()
	tun._destroyCB()
	_ = tun._tunInterface.Close()
	close(tun.OutputStream)
	close(tun.InputStream)
}

func (tun *Tunnel) createTunnelInterface() error {
	CIDRv4 := tun._conf.CIDRv4
	CIDRv6 := tun._conf.CIDRv6
	DeviceName := tun._conf.DeviceName

	c := water.Config{
		DeviceType: water.TUN,
	}
	c.PlatformSpecificParams = water.PlatformSpecificParams{
		Name:    DeviceName,
		Network: []string{CIDRv4, CIDRv6},
	}
	iFace, err := water.New(c)
	if err != nil {
		return fmt.Errorf("failed to create tunnel interface: %v", err)
	}
	tun._tunInterface = iFace
	return nil
}
