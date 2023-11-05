package tunnel

import (
	"github.com/networm6/gopherBox/ctxbox"
	"sync/atomic"
)

func (tun *Tunnel) readFromTunnel() {
	fun := tun._conf.MixinFunc
	packet := make([]byte, tun._conf.BufferSize)
	for ctxbox.Opened(*tun._ctx) {
		num, err := tun._tunInterface.Read(packet)
		tun.incrWrittenBytes(num)
		if err != nil {
			continue
		}
		tun.OutputStream <- fun(packet[:num])
	}
}

func (tun *Tunnel) writeToTunnel() {
	fun := tun._conf.MixinFunc
	for ctxbox.Opened(*tun._ctx) {
		num, err := tun._tunInterface.Write(fun(<-tun.InputStream))
		if err != nil {
			continue
		}
		tun.incrReadBytes(num)
	}
}

func (tun *Tunnel) incrReadBytes(n int) {
	atomic.AddUint64(tun._totalReadBytes, uint64(n))
}

func (tun *Tunnel) incrWrittenBytes(n int) {
	atomic.AddUint64(tun._totalWrittenBytes, uint64(n))
}
