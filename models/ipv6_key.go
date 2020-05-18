package models

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/alex60217101990/types/helpers"
)

type IPv6Key net.IPNet

func (ip *IPv6Key) MarshalJSON() ([]byte, error) {
	if ip == nil {
		ip = &IPv6Key{}
	}
	fmt.Println(ip.IP, ip.Mask)
	m, _ := ip.Mask.Size()
	tmp := fmt.Sprintf("%v/%d", ip.IP, m)
	return json.Marshal(tmp)
}

func (ip *IPv6Key) MarshalYAML() (interface{}, error) {
	if ip == nil {
		ip = &IPv6Key{}
	}
	fmt.Println(ip)
	m, _ := ip.Mask.Size()
	return fmt.Sprintf("%v/%d", ip.IP, m), nil
}

func (ip *IPv6Key) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var tmp string
	if err := unmarshal(&tmp); err != nil {
		return err
	}
	if ip == nil {
		ip = &IPv6Key{}
	}
	var ipnet *net.IPNet
	ipnet, err = helpers.ParseIPFromString(tmp)
	if err != nil {
		return err
	}
	if ipnet == nil {
		return fmt.Errorf("failed parse yaml IPv6, invalid argument: %v", tmp)
	}
	ip.Mask = ipnet.Mask
	ip.IP = ipnet.IP
	return err
}

func (ip *IPv6Key) ParseFromStr(addr string) (err error) {
	if ip == nil {
		ip = &IPv6Key{}
	}
	var ipnet *net.IPNet
	ipnet, err = helpers.ParseIPFromString(addr)
	if err != nil {
		return err
	}
	if ipnet == nil {
		return fmt.Errorf("failed parse IPv6, invalid argument: %v", addr)
	}
	ip.Mask = ipnet.Mask
	ip.IP = ipnet.IP
	return err
}

func (ip *IPv6Key) UnmarshalJSON(data []byte) (err error) {
	var tmp string
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if ip == nil {
		ip = &IPv6Key{}
	}
	var ipnet *net.IPNet
	ipnet, err = helpers.ParseIPFromString(tmp)
	if err != nil {
		return err
	}
	if ipnet == nil {
		return fmt.Errorf("failed parse json IPv6, invalid argument: %v", tmp)
	}
	ip.Mask = ipnet.Mask
	ip.IP = ipnet.IP
	return err
}
