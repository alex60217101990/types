package models

import (
	"encoding/gob"
	"time"

	"github.com/alex60217101990/types/enums"
)

func init() {
	// This type must match exactly what youre going to be using,
	// down to whether or not its a pointer
	gob.Register(&MACBlacklistEvent{})
	gob.Register(&IPBlacklistEvent{})
	gob.Register(&PortBlacklistEvent{})
	gob.Register(&MACBanEvent{})
	gob.Register(&IPBanEvent{})
	gob.Register(&PortBanEvent{})
}

type SidecarEvent interface {
	GetValue() interface{}
}

type MACBlacklistEvent struct {
	MAC    string
	Action enums.NftActionType
}

func (e MACBlacklistEvent) GetValue() interface{} {
	return e
}

type IPBlacklistEvent struct {
	IP     string
	Type   enums.IPType
	Action enums.NftActionType
}

func (e IPBlacklistEvent) GetValue() interface{} {
	return e
}

type PortBlacklistEvent struct {
	Action enums.NftActionType
	Port   PortKey
}

func (e PortBlacklistEvent) GetValue() interface{} {
	return e
}

type MACBanEvent struct {
	MAC      string
	Duration time.Duration
}

func (e MACBanEvent) GetValue() interface{} {
	return e
}

type IPBanEvent struct {
	IP       string
	Type     enums.IPType
	Duration time.Duration
}

func (e IPBanEvent) GetValue() interface{} {
	return e
}

type PortBanEvent struct {
	Port     PortKey
	Duration time.Duration
}

func (e PortBanEvent) GetValue() interface{} {
	return e
}
