package enums

import (
	"encoding/json"
	"fmt"

	"github.com/alex60217101990/types/errors"
	"gopkg.in/yaml.v3"
)

type PortKeyProtocol uint8

const (
	TCP PortKeyProtocol = iota
	UDP
)

var (
	_PortKeyProtocolNameToValue = map[string]PortKeyProtocol{
		"tcp": TCP,
		"TCP": TCP,
		"udp": UDP,
		"UDP": UDP,
	}

	_PortKeyProtocolValueToName = map[PortKeyProtocol]string{
		TCP: "tcp",
		UDP: "udp",
	}
)

func (pkp PortKeyProtocol) MarshalYAML() (interface{}, error) {
	s, ok := _PortKeyProtocolValueToName[pkp]
	if !ok {
		return nil, fmt.Errorf("invalid PortKeyProtocol: %d", pkp)
	}
	return s, nil
}

func (pkp *PortKeyProtocol) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _PortKeyProtocolNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid PortKeyProtocol %q", value.Value)
	}
	*pkp = v
	return nil
}

func (pkp PortKeyProtocol) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(pkp).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PortKeyProtocolValueToName[pkp]
	if !ok {
		return nil, fmt.Errorf("invalid PortKeyProtocol: %d", pkp)
	}
	return json.Marshal(s)
}

func (pkp *PortKeyProtocol) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PortKeyProtocol should be a string, got %s", data)
	}
	v, ok := _PortKeyProtocolNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PortKeyProtocol %q", s)
	}
	*pkp = v
	return nil
}

func (pkp PortKeyProtocol) Val() uint8 {
	return uint8(pkp)
}

// it's for using with flag package
func (pkp *PortKeyProtocol) Set(val string) error {
	if at, ok := _PortKeyProtocolNameToValue[val]; ok {
		*pkp = at
		return nil
	}
	return errors.ErrInvalidPortKeyProtocol(val)
}

func (pkp PortKeyProtocol) String() string {
	return _PortKeyProtocolValueToName[pkp]
}
