package enums

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type ConnServiceType uint8

const (
	Sidecar ConnServiceType = iota
	Agregate
)

var (
	_ConnServiceTypeNameToValue = map[string]ConnServiceType{
		"sidecar":  Sidecar,
		"agregate": Agregate,
	}

	_ConnServiceTypeValueToName = map[ConnServiceType]string{
		Sidecar:  "sidecar",
		Agregate: "agregate",
	}
)

func (cst ConnServiceType) MarshalYAML() (interface{}, error) {
	s, ok := _ConnServiceTypeValueToName[cst]
	if !ok {
		return nil, fmt.Errorf("invalid ConnServiceType: %d", cst)
	}
	return s, nil
}

func (cst *ConnServiceType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _ConnServiceTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid ConnServiceType %q", value.Value)
	}
	*cst = v
	return nil
}

func (cst ConnServiceType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(cst).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ConnServiceTypeValueToName[cst]
	if !ok {
		return nil, fmt.Errorf("invalid ConnServiceType: %d", cst)
	}
	return json.Marshal(s)
}

func (cst *ConnServiceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ConnServiceType should be a string, got %s", data)
	}
	v, ok := _ConnServiceTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ConnServiceType %q", s)
	}
	*cst = v
	return nil
}

func (cst ConnServiceType) Val() uint8 {
	return uint8(cst)
}

// it's for using with flag package
func (cst *ConnServiceType) Set(val string) error {
	if cst == nil {
		var defaultType ConnServiceType
		cst = &defaultType
	}
	if at, ok := _ConnServiceTypeNameToValue[val]; ok {
		*cst = at
		return nil
	}
	return fmt.Errorf("invalid ConnServiceType value: %s", val)
}

func (cst ConnServiceType) String() string {
	return _ConnServiceTypeValueToName[cst]
}
