package enums

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type IPType uint8

const (
	IPv4 IPType = iota
	IPv6
)

var (
	_IPTypeNameToValue = map[string]IPType{
		"ipv4":  IPv4,
		"ipv6":  IPv6,
		"ipv_4": IPv4,
		"ipv_6": IPv6,
		"IPv4":  IPv4,
		"IPv6":  IPv6,
	}

	_IPTypeValueToName = map[IPType]string{
		IPv4: "ipv4",
		IPv6: "ipv6",
	}
)

func (ip IPType) MarshalYAML() (interface{}, error) {
	s, ok := _IPTypeValueToName[ip]
	if !ok {
		return nil, fmt.Errorf("invalid IPType: %d", ip)
	}
	return s, nil
}

func (ip *IPType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _IPTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid IPType %q", value.Value)
	}
	*ip = v
	return nil
}

func (ip IPType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(ip).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _IPTypeValueToName[ip]
	if !ok {
		return nil, fmt.Errorf("invalid IPType: %d", ip)
	}
	return json.Marshal(s)
}

func (ip *IPType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("IPType should be a string, got %s", data)
	}
	v, ok := _IPTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid IPType %q", s)
	}
	*ip = v
	return nil
}

func (ip IPType) Val() uint8 {
	return uint8(ip)
}

// it's for using with flag package
func (ip *IPType) Set(val string) error {
	if ip == nil {
		var defaultType IPType
		ip = &defaultType
	}
	if at, ok := _IPTypeNameToValue[val]; ok {
		*ip = at
		return nil
	}
	return fmt.Errorf("invalid IPType value: %s", val)
}

func (ip IPType) String() string {
	return _IPTypeValueToName[ip]
}
