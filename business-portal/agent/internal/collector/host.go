package collector

import (
	"net"
	"os"
	"runtime"
)

type OSInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type NetworkInfo struct {
	IPAddresses  []string `json:"ipAddresses"`
	MacAddresses []string `json:"macAddresses"`
}

func Hostname() string { h, _ := os.Hostname(); return h }
func OS() OSInfo       { return OSInfo{Name: runtime.GOOS, Version: runtime.GOARCH} }
func Network() NetworkInfo {
	ifs, _ := net.Interfaces()
	var ips, macs []string
	for _, iface := range ifs {
		if iface.HardwareAddr.String() != "" {
			macs = append(macs, iface.HardwareAddr.String())
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ips = append(ips, addr.String())
		}
	}
	return NetworkInfo{IPAddresses: ips, MacAddresses: macs}
}
func Uptime() uint64 { return 0 }
