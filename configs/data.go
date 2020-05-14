package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alex60217101990/goebpf/cgotypes"
	"github.com/alex60217101990/types/enums"
	"github.com/alex60217101990/types/helpers"
	"gopkg.in/yaml.v2"
)

type Configs struct {
	Firewall *Firewall `yaml:"firewall" json:"firewall,omitempty"`
}

type Firewall struct {
	NetIfaceName   *string             `yaml:"net_iface_name,omitempty" json:"net_iface_name,omitempty"`
	ElfFilePath    *string             `yaml:"elf_file_path,omitempty" json:"elf_file_path,omitempty"`
	IPv4BlackList  []string            `yaml:"ipv4_blacklist,omitempty" json:"ipv4_blacklist,omitempty"`
	IPv6BlackList  []string            `yaml:"ipv6_blacklist,omitempty" json:"ipv6_blacklist,omitempty"`
	MacBlacklist   []string            `yaml:"mac_blacklist,omitempty" json:"mac_blacklist,omitempty"`
	PortsBlacklist []*cgotypes.PortKey `yaml:"ports_blacklist,omitempty" json:"ports_blacklist,omitempty"`
}

func (c Configs) PrintTestConfigs(format enums.FormatType, file string) error {
	testConf := &Configs{
		Firewall: &Firewall{
			NetIfaceName: helpers.String("eth0"),
			ElfFilePath:  helpers.String("./tmp/fw.elf"),
			IPv4BlackList: []string{
				"187.162.11.94",
			},
			// PortsBlacklist: []fw.PortKeyConf{
			// 	(*fw.PortKeyConf)(nil).ConvertFromKey(
			// 		fw.PortKey{
			// 			Type:  enums.V4,
			// 			Proto: enums.SourcePort,
			// 			Port:  3128,
			// 		},
			// 	),
			// },
			PortsBlacklist: []*cgotypes.PortKey{
				cgotypes.PortKeyVal(cgotypes.GetPortKey(cgotypes.DestinationPort, cgotypes.UDPPort, 8552)),
				cgotypes.PortKeyVal(cgotypes.GetPortKey(cgotypes.DestinationPort, cgotypes.TCPPort, 3128)),
			},
		},
	}
	var (
		bts []byte
		err error
	)
	switch format {
	case enums.Json:
		bts, err = json.MarshalIndent(testConf, "", "\t")
	case enums.Yaml:
		bts, err = yaml.Marshal(testConf)
	}
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, bts, 0777 /*0664*/)
}

func ReadConfigFile(buff *Configs, format enums.FormatType, path string) (err error) {
	var file []byte
	// var buff Configs
	_, err = os.Stat(path)
	if os.IsNotExist(err) && err != nil {
		path, err = filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}
	}
	file, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	switch format {
	case enums.Json:
		err = json.Unmarshal(file, buff)
	case enums.Yaml:
		err = yaml.Unmarshal(file, buff)
	}
	return err
}
