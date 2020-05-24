package models

import (
	"time"

	"github.com/alex60217101990/types/enums"
)

type MACBlacklistEvent struct {
	MAC    string
	Action enums.NftActionType
}

type IPBlacklistEvent struct {
	IP     string
	Type   enums.IPType
	Action enums.NftActionType
}

type PortBlacklistEvent struct {
	Action enums.NftActionType
	Port   PortKey
}

type MACBanEvent struct {
	MAC      string
	Duration time.Duration
}

type IPBanEvent struct {
	IP       string
	Type     enums.IPType
	Duration time.Duration
}

type PortBanEvent struct {
	Port     PortKey
	Duration time.Duration
}
