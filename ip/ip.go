package ip

import "net"

// ToUint32 将 IP 转为 整型
func ToUint32(s string) uint32 {
	ip := net.ParseIP(s)

	// 无效ip
	if ip == nil {
		return 0
	}

	ip = ip.To4()

	var out uint32

	out += uint32(ip[0]) << 24
	out += uint32(ip[1]) << 16
	out += uint32(ip[2]) << 8
	out += uint32(ip[3])

	return out
}

// ToString 将 整型 转为 ipv4(长度 11)
func ToString(n uint32) string {
	ip := make(net.IP, net.IPv4len)

	ip[0] = byte((n >> 24) & 0xFF)
	ip[1] = byte((n >> 16) & 0xFF)
	ip[2] = byte((n >> 8) & 0xFF)
	ip[3] = byte(n & 0xFF)
	return ip.String()
}

// GetLocalAddr 获取本地地址，只限 ipv4
func GetLocalAddr() ([]string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}

	var out []string

	for _, addr := range addrs {
		if n, ok := addr.(*net.IPNet); ok && !n.IP.IsLoopback() {
			if n.IP.To4() != nil /** || n.IP.To16() != nil */ {
				out = append(out, n.IP.String())
			}
		}
	}

	return out, nil
}
