package enums

import (
	"encoding/json"
	"fmt"

	"github.com/alex60217101990/types/errors"
	"gopkg.in/yaml.v3"
)

type NftActionType int

const (
	AddAction NftActionType = iota
	DeleteAction
)

var (
	_NftActionTypeNameToValue = map[string]NftActionType{
		"Add":    AddAction,
		"Delete": DeleteAction,
		"add":    AddAction,
		"delete": DeleteAction,
	}

	_NftActionTypeValueToName = map[NftActionType]string{
		AddAction:    "add",
		DeleteAction: "delete",
	}
)

func (a NftActionType) MarshalYAML() (interface{}, error) {
	s, ok := _NftActionTypeValueToName[a]
	if !ok {
		return nil, fmt.Errorf("invalid NftActionType: %d", a)
	}
	return s, nil
}

func (a *NftActionType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _NftActionTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid NftActionType %q", value.Value)
	}
	*a = v
	return nil
}

func (a NftActionType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(a).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _NftActionTypeValueToName[a]
	if !ok {
		return nil, fmt.Errorf("invalid NftActionType: %d", a)
	}
	return json.Marshal(s)
}

func (a *NftActionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NftActionType should be a string, got %s", data)
	}
	v, ok := _NftActionTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid NftActionType %q", s)
	}
	*a = v
	return nil
}

func (a NftActionType) Val() int {
	return int(a)
}

// it's for using with flag package
func (a *NftActionType) Set(val string) error {
	if at, ok := _NftActionTypeNameToValue[val]; ok {
		*a = at
		return nil
	}
	return errors.ErrInvalidActionType(val)
}

func (a NftActionType) String() string {
	return _NftActionTypeValueToName[a]
}
