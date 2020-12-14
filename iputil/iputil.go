package iputil

import (
	"net"
)

// IsPublic 判断ip是否是公网ip
func IsPublic(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}

	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		case ip4[0] == 100 && ip4[1] >= 64 && ip4[1] <= 127:
			return false
		case ip4[0] == 169 && ip4[1] == 254:
			return false
		default:
			return true
		}
	}

	return false
}

// MaskSize 将掩码转换为数字，入参格式 255.255.128.0
func MaskSize(mask string) int {
	ipMask := net.IPMask(net.ParseIP(mask).To4())
	ones, bits := ipMask.Size()
	if bits == 0 {
		return -1
	}
	return ones
}

// ParseCidr 解析出段下所有ip，入参格式 192.168.1.1/24
func ParseIPRange(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
