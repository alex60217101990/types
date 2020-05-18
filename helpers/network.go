package helpers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() (ip *string, err error) {
	var addr string
	func() {
		conn, err1 := net.Dial("udp", "8.8.8.8:80")
		if err1 != nil {
			err = err1
		}
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
			err = conn.Close()
		}()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		addr = localAddr.IP.To4().String()
	}()
	if err != nil || len(addr) == 0 {
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = r.(error)
				}
			}()
			addrs, err2 := net.InterfaceAddrs()
			if err2 != nil {
				err = err2
			}
			if addrs != nil {
				if ipnet, ok := addrs[0].(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						addr = ipnet.IP.String()
					}
				}
			} else {
				err = fmt.Errorf("can't get net addresses")
			}
		}()
	}
	if err != nil {
		return nil, err
	}
	return String(addr), nil
}

func ParseIPFromString(s string) (ipnet *net.IPNet, err error) {
	// Check if given address is CIDR
	if strings.Contains(s, "/") {
		_, ipnet, err = net.ParseCIDR(s)
	} else {
		if strings.Contains(s, ":") {
			// IPv6
			_, ipnet, err = net.ParseCIDR(s + "/128")
		} else {
			// IPv4
			_, ipnet, err = net.ParseCIDR(s + "/32")
		}
	}
	return ipnet, err
}

func IP2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func Int2IP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
