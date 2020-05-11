package enums

import (
	"encoding/json"
	"fmt"

	"github.com/alex60217101990/types/errors"
	"gopkg.in/yaml.v3"
)

type FormatType uint8

const (
	Yaml FormatType = iota
	Json
	Xml
	Env
)

var (
	_FormatTypeNameToValue = map[string]FormatType{
		"yaml": Yaml,
		"json": Json,
		"xml":  Xml,
		"env":  Env,
	}

	_FormatTypeValueToName = map[FormatType]string{
		Yaml: "yaml",
		Json: "json",
		Xml:  "xml",
		Env:  "env",
	}
)

func (f FormatType) MarshalYAML() (interface{}, error) {
	s, ok := _FormatTypeValueToName[f]
	if !ok {
		return nil, fmt.Errorf("invalid FormatType: %d", f)
	}
	return s, nil
}

func (f *FormatType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _FormatTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid FormatType %q", value.Value)
	}
	*f = v
	return nil
}

func (f FormatType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(f).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _FormatTypeValueToName[f]
	if !ok {
		return nil, fmt.Errorf("invalid FormatType: %d", f)
	}
	return json.Marshal(s)
}

func (f *FormatType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("FormatType should be a string, got %s", data)
	}
	v, ok := _FormatTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid FormatType %q", s)
	}
	*f = v
	return nil
}

func (f FormatType) Val() uint8 {
	return uint8(f)
}

// it's for using with flag package
func (f *FormatType) Set(val string) error {
	if at, ok := _FormatTypeNameToValue[val]; ok {
		*f = at
		return nil
	}
	return errors.ErrInvalidFormatType(val)
}

func (f FormatType) String() string {
	return _FormatTypeValueToName[f]
}
