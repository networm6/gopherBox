package water

import (
	"io"
	"reflect"
)

// MultiQueue(Linux kernel > 3.8): With MultiQueue enabled, user should hold multiple
// interfaces to send/receive packet in parallel.
// Kernel document about MultiQueue: https://www.kernel.org/doc/Documentation/networking/tuntap.txt

// copy from https://github.com/net-byte/water

type Interface struct {
	io.ReadWriteCloser
	name string
}

// Config defines parameters required to create a TUN/TAP interface. It's only
// used when the device is initialized. A zero-value Config is a valid
// configuration.
type Config struct {

	// PlatformSpecificParams defines parameters that differ on different
	// platforms. See comments for the type for more details.
	PlatformSpecificParams
}

func defaultConfig() Config {
	return Config{
		PlatformSpecificParams: defaultPlatformSpecificParams(),
	}
}

var zeroConfig Config

// New creates a new TUN/TAP interface using config.
func New(config Config) (ifce *Interface, err error) {
	if reflect.DeepEqual(config, zeroConfig) {
		config = defaultConfig()
	}
	if reflect.DeepEqual(config.PlatformSpecificParams, zeroConfig.PlatformSpecificParams) {
		config.PlatformSpecificParams = defaultPlatformSpecificParams()
	}

	return openDev(config)

}

// Name returns the interface name of ifce, e.g. tun0, tap1, tun0, etc..
func (ifce *Interface) Name() string {
	return ifce.name
}
