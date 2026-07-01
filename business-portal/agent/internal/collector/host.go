package collector
import("github.com/shirou/gopsutil/v3/host";"github.com/shirou/gopsutil/v3/net";"os")
type OSInfo struct{Name string `json:"name"`; Version string `json:"version"`}; type NetworkInfo struct{IPAddresses []string `json:"ipAddresses"`; MacAddresses []string `json:"macAddresses"`}
func Hostname()string{h,_:=os.Hostname(); return h}
func OS()(OSInfo){i,_:=host.Info(); return OSInfo{Name:i.Platform,Version:i.PlatformVersion}}
func Network()NetworkInfo{ifs,_:=net.Interfaces(); var ips,macs []string; for _,i:= range ifs{if i.HardwareAddr!=""{macs=append(macs,i.HardwareAddr)}; for _,a:= range i.Addrs{ips=append(ips,a.Addr)}}; return NetworkInfo{ips,macs}}
func Uptime()uint64{u,_:=host.Uptime(); return u}
