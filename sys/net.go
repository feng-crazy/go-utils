package sys

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)

	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// ignore docker and warden bridge
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}

	return ips, nil
}

func IsIntranet(ipStr string) bool {
	if strings.HasPrefix(ipStr, "10.") {
		return true
	}

	// for didi
	if strings.HasPrefix(ipStr, "100.") {
		return true
	}

	if strings.HasPrefix(ipStr, "192.168.") {
		return true
	}

	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

// ${sn}-${hostname}-${ip}
func LocalHostIdent() string {
	sn, _ := CmdOutTrim("/bin/bash", "-c", "dmidecode -s system-serial-number")
	if sn != "" {
		arr := strings.Fields(sn)
		sn = arr[len(arr)-1]
	} else {
		sn = "nil"
	}

	name, _ := CmdOutTrim("hostname")

	ips, _ := IntranetIP()
	ip := ""
	if ips != nil && len(ips) > 0 {
		ip = ips[0]
	}

	return fmt.Sprintf("%s-%s-%s", sn, name, ip)
}

func GetOutboundIpaddr() string {
	conn, err := net.Dial("udp4", "1.2.3.4:56")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	if ip, _, err := net.SplitHostPort(localAddr); err != nil {
		return ""
	} else {
		return ip
	}
}

// IsPublicIP 判断ip是否是公网IP
func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// GetLanIPs 获取局域网本地IP地址列表
func GetLanIPs() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok &&
			!ipnet.IP.IsLoopback() &&
			!ipnet.IP.IsLinkLocalMulticast() &&
			!ipnet.IP.IsLinkLocalUnicast() &&
			IsIPv4(ipnet.IP) &&
			!IsPublicIP(ipnet.IP) {

			ips = append(ips, ipnet.IP.String())
		}
	}

	return
}

// ParseCIDRs parses a list of cidrs and return error if any is invalid.
// order is maintained
func ParseCIDRs(cidrsString []string) ([]*net.IPNet, error) {
	cidrs := make([]*net.IPNet, 0, len(cidrsString))
	for _, cidrString := range cidrsString {
		_, cidr, err := net.ParseCIDR(cidrString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse cidr value:%q with error:%v", cidrString, err)
		}
		cidrs = append(cidrs, cidr)
	}
	return cidrs, nil
}

// IsDualStackIPs returns if a slice of ips is:
// - all are valid ips
// - at least one ip from each family (v4 or v6)
func IsDualStackIPs(ips []net.IP) (bool, error) {
	v4Found := false
	v6Found := false
	for _, ip := range ips {
		if ip == nil {
			return false, fmt.Errorf("ip %v is invalid", ip)
		}

		if v4Found && v6Found {
			continue
		}

		if IsIPv6(ip) {
			v6Found = true
			continue
		}

		v4Found = true
	}

	return v4Found && v6Found, nil
}

// IsDualStackIPStrings returns if
// - all are valid ips
// - at least one ip from each family (v4 or v6)
func IsDualStackIPStrings(ips []string) (bool, error) {
	parsedIPs := make([]net.IP, 0, len(ips))
	for _, ip := range ips {
		parsedIP := net.ParseIP(ip)
		parsedIPs = append(parsedIPs, parsedIP)
	}
	return IsDualStackIPs(parsedIPs)
}

// IsDualStackCIDRs returns if
// - all are valid cidrs
// - at least one cidr from each family (v4 or v6)
func IsDualStackCIDRs(cidrs []*net.IPNet) (bool, error) {
	v4Found := false
	v6Found := false
	for _, cidr := range cidrs {
		if cidr == nil {
			return false, fmt.Errorf("cidr %v is invalid", cidr)
		}

		if v4Found && v6Found {
			continue
		}

		if IsIPv6(cidr.IP) {
			v6Found = true
			continue
		}
		v4Found = true
	}

	return v4Found && v6Found, nil
}

// IsDualStackCIDRStrings returns if
// - all are valid cidrs
// - at least one cidr from each family (v4 or v6)
func IsDualStackCIDRStrings(cidrs []string) (bool, error) {
	parsedCIDRs, err := ParseCIDRs(cidrs)
	if err != nil {
		return false, err
	}
	return IsDualStackCIDRs(parsedCIDRs)
}

// IsIPv6 returns if netIP is IPv6.
func IsIPv6(netIP net.IP) bool {
	return netIP != nil && netIP.To4() == nil
}

// IsIPv6String returns if ip is IPv6.
func IsIPv6String(ip string) bool {
	netIP := net.ParseIP(ip)
	return IsIPv6(netIP)
}

// IsIPv6CIDRString returns if cidr is IPv6.
// This assumes cidr is a valid CIDR.
func IsIPv6CIDRString(cidr string) bool {
	ip, _, _ := net.ParseCIDR(cidr)
	return IsIPv6(ip)
}

// IsIPv6CIDR returns if a cidr is ipv6
func IsIPv6CIDR(cidr *net.IPNet) bool {
	ip := cidr.IP
	return IsIPv6(ip)
}

// IsIPv4 returns if netIP is IPv4.
func IsIPv4(netIP net.IP) bool {
	return netIP != nil && netIP.To4() != nil
}

// IsIPv4String returns if ip is IPv4.
func IsIPv4String(ip string) bool {
	netIP := net.ParseIP(ip)
	return IsIPv4(netIP)
}

// IsIPv4CIDR returns if a cidr is ipv4
func IsIPv4CIDR(cidr *net.IPNet) bool {
	ip := cidr.IP
	return IsIPv4(ip)
}

// IsIPv4CIDRString returns if cidr is IPv4.
// This assumes cidr is a valid CIDR.
func IsIPv4CIDRString(cidr string) bool {
	ip, _, _ := net.ParseCIDR(cidr)
	return IsIPv4(ip)
}

// ParsePort parses a string representing an IP port.  If the string is not a
// valid port number, this returns an error.
func ParsePort(port string, allowZero bool) (int, error) {
	portInt, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return 0, err
	}
	if portInt == 0 && !allowZero {
		return 0, errors.New("0 is not a valid port number")
	}
	return int(portInt), nil
}

// BigForIP creates a big.Int based on the provided net.IP
func BigForIP(ip net.IP) *big.Int {
	// NOTE: Convert to 16-byte representation so we can
	// handle v4 and v6 values the same way.
	return big.NewInt(0).SetBytes(ip.To16())
}

// AddIPOffset adds the provided integer offset to a base big.Int representing a net.IP
// NOTE: If you started with a v4 address and overflow it, you get a v6 result.
func AddIPOffset(base *big.Int, offset int) net.IP {
	r := big.NewInt(0).Add(base, big.NewInt(int64(offset))).Bytes()
	r = append(make([]byte, 16), r...)
	return net.IP(r[len(r)-16:])
}

// RangeSize returns the size of a range in valid addresses.
// returns the size of the subnet (or math.MaxInt64 if the range size would overflow int64)
func RangeSize(subnet *net.IPNet) int64 {
	ones, bits := subnet.Mask.Size()
	if bits == 32 && (bits-ones) >= 31 || bits == 128 && (bits-ones) >= 127 {
		return 0
	}
	// this checks that we are not overflowing an int64
	if bits-ones >= 63 {
		return math.MaxInt64
	}
	return int64(1) << uint(bits-ones)
}

// GetIndexedIP returns a net.IP that is subnet.IP + index in the contiguous IP space.
func GetIndexedIP(subnet *net.IPNet, index int) (net.IP, error) {
	ip := AddIPOffset(BigForIP(subnet.IP), index)
	if !subnet.Contains(ip) {
		return nil, fmt.Errorf("can't generate IP with index %d from subnet. subnet too small. subnet: %q", index, subnet)
	}
	return ip, nil
}

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// 获取本机网卡IP
func GetLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	err = errors.New("没有找到网卡IP")
	return
}
