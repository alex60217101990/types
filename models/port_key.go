package models

import (
	"encoding/json"
	"fmt"

	"github.com/alex60217101990/types/enums"
)

type PortKey struct {
	Type  enums.PortKeyType
	Proto enums.PortKeyProtocol
	Port  uint16
}

func (p *PortKey) MarshalJSON() ([]byte, error) {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint16 `yaml:"port" json:"port"`
	}
	if p == nil {
		p = &PortKey{}
	}
	return json.Marshal(alias{
		Type:  p.Type.String(),
		Proto: p.Proto.String(),
		Port:  p.Port,
	})
}

func (p *PortKey) UnmarshalJSON(data []byte) (err error) {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint16 `yaml:"port" json:"port"`
	}
	var tmp alias
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if p == nil {
		p = &PortKey{}
	}
	err = p.Type.Set(tmp.Type)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortType: %v", tmp.Type, err)
	}
	err = p.Proto.Set(tmp.Proto)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortProtocol: %v", tmp.Proto, err)
	}
	p.Port = tmp.Port

	return nil
}

func (p *PortKey) MarshalYAML() (interface{}, error) {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint16 `yaml:"port" json:"port"`
	}
	if p == nil {
		p = &PortKey{}
	}
	return alias{
		Type:  p.Type.String(),
		Proto: p.Proto.String(),
		Port:  p.Port,
	}, nil
}

func (p *PortKey) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint16 `yaml:"port" json:"port"`
	}
	var tmp alias
	if err := unmarshal(&tmp); err != nil {
		return err
	}
	if p == nil {
		p = &PortKey{}
	}
	err := p.Type.Set(tmp.Type)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortType: %v", tmp.Type, err)
	}
	err = p.Proto.Set(tmp.Proto)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortProtocol: %v", tmp.Proto, err)
	}
	p.Port = tmp.Port

	return nil
}
