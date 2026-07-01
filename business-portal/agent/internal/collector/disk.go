package collector
import "github.com/shirou/gopsutil/v3/disk"
func DiskPercent()float64{parts,_:=disk.Partitions(false); var max float64; for _,p:=range parts{u,_:=disk.Usage(p.Mountpoint); if u!=nil&&u.UsedPercent>max{max=u.UsedPercent}}; return max}
