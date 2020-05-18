package enums

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type PortKeyType uint8

const (
	DestinationPort PortKeyType = iota
	SourcePort
)

var (
	_PortKeyTypeNameToValue = map[string]PortKeyType{
		"source_port":      SourcePort,
		"destination_port": DestinationPort,
	}

	_PortKeyTypeValueToName = map[PortKeyType]string{
		SourcePort:      "source_port",
		DestinationPort: "destination_port",
	}
)

func (p PortKeyType) MarshalYAML() (interface{}, error) {
	s, ok := _PortKeyTypeValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortKeyType: %d", p)
	}
	return s, nil
}

func (p *PortKeyType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _PortKeyTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid PortKeyType %q", value.Value)
	}
	*p = v
	return nil
}

func (p PortKeyType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PortKeyTypeValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortKeyType: %d", p)
	}
	return json.Marshal(s)
}

func (p *PortKeyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PortKeyType should be a string, got %s", data)
	}
	v, ok := _PortKeyTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PortKeyType %q", s)
	}
	*p = v
	return nil
}

func (p PortKeyType) Val() uint8 {
	return uint8(p)
}

// it's for using with flag package
func (p *PortKeyType) Set(val string) error {
	if p == nil {
		var defaultType PortKeyType
		p = &defaultType
	}
	if at, ok := _PortKeyTypeNameToValue[val]; ok {
		*p = at
		return nil
	}
	return fmt.Errorf("invalid PortKeyType value: %s", val)
}

func (p PortKeyType) String() string {
	return _PortKeyTypeValueToName[p]
}
