package models

import "github.com/alex60217101990/types/enums"

type TCPConnMeta struct {
	ServiceType enums.ConnServiceType
	ServiceID   string
	ServiceAddr string
}
