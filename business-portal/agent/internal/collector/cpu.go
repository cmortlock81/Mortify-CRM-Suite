package collector
import("time";"github.com/shirou/gopsutil/v3/cpu")
func CPUPercent()float64{v,_:=cpu.Percent(time.Second,false); if len(v)>0{return v[0]}; return 0}
