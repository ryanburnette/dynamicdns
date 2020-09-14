package ipify

import (
	"github.com/rdegges/go-ipify"
)

// GetIP gets your current public IP address via Ipify
func GetIP() string {
	ip, err := ipify.GetIp()
	if err != nil {
		panic(err)
	}
	return ip
}
