package netbox

import (
	"log"
	"net"
	"strings"
)

// GetInterface returns the name of interface
func GetInterface() (name string) {
	interfaces := getAllInterfaces()
	if len(interfaces) == 0 {
		return ""
	}
	addresses, _ := interfaces[0].Addrs()
	for _, addr := range addresses {
		ip, ok := addr.(*net.IPNet)
		if ok && ip.IP.To4() != nil && !ip.IP.IsLoopback() {
			name = interfaces[0].Name
			break
		}
	}
	return name
}

// getAllInterfaces returns all interfaces
func getAllInterfaces() []net.Interface {
	iFaceList, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return nil
	}

	var outInterfaces []net.Interface
	for _, iFace := range iFaceList {
		if iFace.Flags&net.FlagLoopback == 0 && iFace.Flags&net.FlagUp == 1 && isPhysicalInterface(iFace.Name) {
			netAddrList, _ := iFace.Addrs()
			if len(netAddrList) > 0 {
				outInterfaces = append(outInterfaces, iFace)
			}
		}
	}
	return outInterfaces
}

// isPhysicalInterface returns true if the interface is physical
func isPhysicalInterface(addr string) bool {
	prefixArray := []string{"ens", "enp", "enx", "eno", "eth", "en0", "wlan", "wlp", "wlo", "wlx", "wifi0", "lan0"}
	for _, pref := range prefixArray {
		if strings.HasPrefix(strings.ToLower(addr), pref) {
			return true
		}
	}
	return false
}

// LookupIP Lookup IP address of the given hostname
func LookupIP(domain string) net.IP {
	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		log.Println(err)
		return nil
	}
	return ips[0]
}

// DiscoverGateway returns the local gateway IP address
func DiscoverGateway(ipv4 bool) string {
	var ip net.IP
	var err error
	if ipv4 {
		ip, err = DiscoverGatewayIPv4()
	} else {
		ip, err = DiscoverGatewayIPv6()
	}
	if err != nil {
		log.Println(err)
		return ""
	}
	return ip.String()
}

// LookupServerAddrIP returns the IP of server address
func LookupServerAddrIP(serverAddr string) net.IP {
	host, _, err := net.SplitHostPort(serverAddr)
	if err != nil {
		log.Panic("error server address")
		return nil
	}
	ip := LookupIP(host)
	return ip
}

// GetSrcKey returns the source key of the packet
func GetSrcKey(packet []byte) string {
	key := ""
	if len(packet) >= 20 && IsIPv4(packet) {
		key = GetIPv4Src(packet).To4().String()
	} else if len(packet) >= 40 && IsIPv6(packet) {
		key = GetIPv6Src(packet).To16().String()
	}
	return key
}

// GetDstKey returns the destination key of the packets
func GetDstKey(packet []byte) string {
	key := ""
	if IsIPv4(packet) && len(packet) >= 20 {
		key = GetIPv4Dst(packet).To4().String()
	} else if IsIPv6(packet) && len(packet) >= 40 {
		key = GetIPv6Dst(packet).To16().String()
	}
	return key
}

// IsIPv4 returns true if the packet is IPv4s
func IsIPv4(packet []byte) bool {
	flag := packet[0] >> 4
	return flag == 4
}

// IsIPv6 returns true if the packet is IPv6s
func IsIPv6(packet []byte) bool {
	flag := packet[0] >> 4
	return flag == 6
}

// GetIPv4Src returns the IPv4 source address of the packet
func GetIPv4Src(packet []byte) net.IP {
	return net.IPv4(packet[12], packet[13], packet[14], packet[15])
}

// GetIPv4Dst returns the IPv4 destination address of the packet
func GetIPv4Dst(packet []byte) net.IP {
	return net.IPv4(packet[16], packet[17], packet[18], packet[19])
}

// GetIPv6Src returns the IPv6 source address of the packet
func GetIPv6Src(packet []byte) net.IP {
	return packet[8:24]
}

// GetIPv6Dst returns the IPv6 destination address of the packet
func GetIPv6Dst(packet []byte) net.IP {
	return packet[24:40]
}
