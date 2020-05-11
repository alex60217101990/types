package enums

import (
	"encoding/json"
	"fmt"

	"github.com/alex60217101990/types/errors"
	"gopkg.in/yaml.v3"
)

type PortType int32

const (
	SourcePort PortType = iota
	DestinationPort
)

var (
	_PortTypeNameToValue = map[string]PortType{
		"source_port":      SourcePort,
		"destination_port": DestinationPort,
	}

	_PortTypeValueToName = map[PortType]string{
		SourcePort:      "source_port",
		DestinationPort: "destination_port",
	}
)

func (p PortType) MarshalYAML() (interface{}, error) {
	s, ok := _PortTypeValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortType: %d", p)
	}
	return s, nil
}

func (p *PortType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _PortTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid PortType %q", value.Value)
	}
	*p = v
	return nil
}

func (p PortType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PortTypeValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortType: %d", p)
	}
	return json.Marshal(s)
}

func (p *PortType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PortType should be a string, got %s", data)
	}
	v, ok := _PortTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PortType %q", s)
	}
	*p = v
	return nil
}

func (p PortType) Val() int32 {
	return int32(p)
}

// it's for using with flag package
func (p *PortType) Set(val string) error {
	if p == nil {
		var defaultType PortType
		p = &defaultType
	}
	if at, ok := _PortTypeNameToValue[val]; ok {
		*p = at
		return nil
	}
	return errors.ErrInvalidPortType(val)
}

func (p PortType) String() string {
	return _PortTypeValueToName[p]
}
