package errors

import (
	"fmt"

	"github.com/juju/errors"
)

var (
	ErrInvalidFormatType = func(val string) error { return fmt.Errorf("invalid format type: %v", val) }
	ErrInvalidPortType   = func(val string) error { return fmt.Errorf("invalid port type: %v", val) }
	// ebpf loader
	ErrBpfSystemConfigNil = errors.New("system configs not set")
	ErrXdpMapKeyDataType  = errors.New("invalid xdp map event data type")
	ErrBpfInitFatal       = func(val string) error { return fmt.Errorf("ebpf init fatal: %v", val) }
	ErrLoadingFwElfFile   = func(val string) error { return fmt.Errorf("Loading firewall elf file failed: %v", val) }
	ErrMapNotFound        = func(val string) error { return fmt.Errorf("eBPF map '%s' not found", val) }
	ErrProgNotFound       = func(val string) error { return fmt.Errorf("Program '%s' not found", val) }
	ErrLoadXdpProg        = func(val ...string) error { return fmt.Errorf("Loading %s xdp program failed: %v", val...) }
	ErrAttachXdpProg      = func(val ...string) error {
		return fmt.Errorf("Attach %s xdp program to iface: %s failed: %v", val...)
	}
	ErrParseMACAddr = func(val ...string) error {
		return fmt.Errorf("parse MAC addr: %s, from configs blaclist error: %v", val...)
	}
	ErrLoadFwBlacklists = func(val string) error { return fmt.Errorf("Load firewall blacklists: %v", val) }
	ErrAttachIfaceToXdpProg = func(val ...string) error { return fmt.Errorf("can't attach interface: %s to xdp program: %s", val...)
)
