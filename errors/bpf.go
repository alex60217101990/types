package errors

import (
	"fmt"
	"reflect"

	"github.com/juju/errors"
)

var (
	ErrInvalidFormatType      = func(val interface{}) error { return fmt.Errorf("invalid format type: %v", val) }
	ErrInvalidPortType        = func(val interface{}) error { return fmt.Errorf("invalid port type: %v", val) }
	ErrInvalidPortKeyProtocol = func(val interface{}) error { return fmt.Errorf("invalid port key protocol type: %v", val) }
	// ebpf loader
	ErrBpfSystemConfigNil = errors.New("system configs not set")
	ErrXdpMapKeyDataType  = errors.New("invalid xdp map event data type")
	ErrBpfInitFatal       = func(val interface{}) error { return fmt.Errorf("ebpf init fatal: %v", val) }
	ErrLoadingFwElfFile   = func(val interface{}) error { return fmt.Errorf("Loading firewall elf file failed: %v", val) }
	ErrMapNotFound        = func(val interface{}) error { return fmt.Errorf("eBPF map '%v' not found", val) }
	ErrProgNotFound       = func(val interface{}) error { return fmt.Errorf("Program '%v' not found", val) }
	ErrLoadXdpProg        = func(val ...interface{}) error { return fmt.Errorf("Loading %v xdp program failed: %v", val...) }
	ErrAttachXdpProg      = func(val ...interface{}) error {
		return fmt.Errorf("Attach %v xdp program to iface: %v failed: %v", val...)
	}
	ErrParseMACAddr = func(val ...interface{}) error {
		return fmt.Errorf("parse MAC addr: %v, from configs blaclist error: %v", val...)
	}
	ErrLoadFwBlacklists     = func(val interface{}) error { return fmt.Errorf("Load firewall blacklists: %v", val) }
	ErrAttachIfaceToXdpProg = func(val ...interface{}) error {
		return fmt.Errorf("can't attach interface: %v to xdp program: %s", val...)
	}
	ErrParseEnumYamlField = func(val ...interface{}) error {
		return fmt.Errorf("failed to parse '%s' to enums.%s: %v", val[0], reflect.TypeOf(val[1]).Elem().Name(), val[2])
	}
	//ErrParseIPTypeYamlField = func(val ...interface{}) error { return fmt.Errorf("failed to parse '%s' to enums.IpType: %v", val...) }
)
