package models

import (
	"time"

	"github.com/alex60217101990/types/enums"
)

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
