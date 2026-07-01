package collector
import "github.com/shirou/gopsutil/v3/mem"
func MemoryPercent()float64{v,_:=mem.VirtualMemory(); if v==nil{return 0}; return v.UsedPercent}
