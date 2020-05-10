package enums

import (
	"encoding/json"
	"fmt"

	"github.com/alex60217101990/types/errors"
	"gopkg.in/yaml.v3"
)

type IpType int32

const (
	V4 IpType = iota
	V6
)

var (
	_IpTypeNameToValue = map[string]IpType{
		"IPv4": V4,
		"IPv6": V6,
	}

	_IpTypeValueToName = map[IpType]string{
		V4: "IPv4",
		V6: "IPv6",
	}
)

func (ip IpType) MarshalYAML() (interface{}, error) {
	s, ok := _IpTypeValueToName[ip]
	if !ok {
		return nil, fmt.Errorf("invalid IpType: %d", ip)
	}
	return s, nil
}

func (ip *IpType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _IpTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid IpType %q", value.Value)
	}
	*ip = v
	return nil
}

func (ip IpType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(ip).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _IpTypeValueToName[ip]
	if !ok {
		return nil, fmt.Errorf("invalid IpType: %d", ip)
	}
	return json.Marshal(s)
}

func (ip *IpType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("IpType should be a string, got %s", data)
	}
	v, ok := _IpTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid IpType %q", s)
	}
	*ip = v
	return nil
}

func (ip IpType) Val() int32 {
	return int32(ip)
}

// it's for using with flag package
func (ip *IpType) Set(val string) error {
	if at, ok := _IpTypeNameToValue[val]; ok {
		*ip = at
		return nil
	}
	return errors.ErrInvalidIPVersionType(val)
}

func (ip IpType) String() string {
	return _IpTypeValueToName[ip]
}
