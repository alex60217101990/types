package fw

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/alex60217101990/types/consts"
	"github.com/alex60217101990/types/enums"
)

type LpmV4Key struct {
	Prefixlen uint32
	Address   [4]uint8
}

func (l *LpmV4Key) ParseFromSrt(ipStr string) (err error) {
	var ipnet *net.IPNet
	// Check if given address is CIDR
	if strings.Contains(ipStr, "/") {
		_, ipnet, err = net.ParseCIDR(ipStr)
	} else {
		// IPv4
		_, ipnet, err = net.ParseCIDR(ipStr + "/32")
	}
	if err != nil {
		log.Println(err)
		return err
	}
	maskVal, _ := ipnet.Mask.Size()
	*l = LpmV4Key{
		Prefixlen: uint32(maskVal),
	}

	copy(l.Address[:], []byte(ipnet.IP.To4()))
	log.Printf("%s => %v\n", []byte(ipnet.IP.To4().String()), l.Address)
	return err
}

func ubtoa(dst []byte, start int, v byte) int {
	if v < 10 {
		dst[start] = v + '0'
		return 1
	} else if v < 100 {
		dst[start+1] = v%10 + '0'
		dst[start] = v/10 + '0'
		return 2
	}

	dst[start+2] = v%10 + '0'
	dst[start+1] = (v/10)%10 + '0'
	dst[start] = v/100 + '0'
	return 3
}

func (l LpmV4Key) String() string {
	const maxIPv4StringLen = len("255.255.255.255")
	b := make([]byte, maxIPv4StringLen)

	n := ubtoa(b, 0, l.Address[0])
	b[n] = '.'
	n++

	n += ubtoa(b, n, l.Address[1])
	b[n] = '.'
	n++

	n += ubtoa(b, n, l.Address[2])
	b[n] = '.'
	n++

	n += ubtoa(b, n, l.Address[3])
	if l.Prefixlen > 0 {
		return fmt.Sprintf("%s/%d", string(b[:n]), l.Prefixlen)
	}
	return string(b[:n])
}

type LpmV6Key struct {
	Prefixlen uint32
	Address   [16]uint8
}

func (l *LpmV6Key) ParseFromSrt(ipStr string) (err error) {
	var ipnet *net.IPNet
	// Check if given address is CIDR
	if strings.Contains(ipStr, "/") {
		_, ipnet, err = net.ParseCIDR(ipStr)
	} else {
		// IPv6
		_, ipnet, err = net.ParseCIDR(ipStr + "/128")
	}
	if err != nil {
		log.Println(err)
		return err
	}
	maskVal, _ := ipnet.Mask.Size()
	*l = LpmV6Key{
		Prefixlen: uint32(maskVal),
	}

	copy(l.Address[:], []byte(ipnet.IP.To16()))
	// log.Printf("%s => %v\n", []byte(ipnet.IP.To16().String()), l.Address)
	return err
}

func hexString(b []byte) string {
	s := make([]byte, len(b)*2)
	for i, tn := range b {
		s[i*2], s[i*2+1] = consts.HexDigit[tn>>4], consts.HexDigit[tn&0xf]
	}
	return string(s)
}

func appendHex(dst []byte, i uint32) []byte {
	if i == 0 {
		return append(dst, '0')
	}
	for j := 7; j >= 0; j-- {
		v := i >> uint(j*4)
		if v > 0 {
			dst = append(dst, consts.HexDigit[v&0xf])
		}
	}
	return dst
}

func (l LpmV6Key) String() string {
	if len(l.Address) != net.IPv6len {
		return "?" + hexString(l.Address[:])
	}
	// Find longest run of zeros.
	e0 := -1
	e1 := -1
	for i := 0; i < net.IPv6len; i += 2 {
		j := i
		for j < net.IPv6len && l.Address[j] == 0 && l.Address[j+1] == 0 {
			j += 2
		}
		if j > i && j-i > e1-e0 {
			e0 = i
			e1 = j
			i = j
		}
	}
	// The symbol "::" MUST NOT be used to shorten just one 16 bit 0 field.
	if e1-e0 <= 2 {
		e0 = -1
		e1 = -1
	}

	const maxLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	b := make([]byte, 0, maxLen)

	// Print with possible :: in place of run of zeros
	for i := 0; i < net.IPv6len; i += 2 {
		if i == e0 {
			b = append(b, ':', ':')
			i = e1
			if i >= net.IPv6len {
				break
			}
		} else if i > 0 {
			b = append(b, ':')
		}
		b = appendHex(b, (uint32(l.Address[i])<<8)|uint32(l.Address[i+1]))
	}
	if l.Prefixlen > 0 {
		return fmt.Sprintf("%s/%d", string(b), l.Prefixlen)
	}
	return string(b)
}

type PortKey struct {
	Type_ enums.IpType
	Proto enums.PortType
	Port  uint32
}
