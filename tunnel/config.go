package tunnel

type TunConfig struct {
	DeviceName string `json:"device_name"`
	MTU        int    `json:"mtu"`

	BufferSize int                 `json:"buffer_size"`
	MixinFunc  func([]byte) []byte `json:"mixin_func"`

	CIDRv4 string
	CIDRv6 string
}
