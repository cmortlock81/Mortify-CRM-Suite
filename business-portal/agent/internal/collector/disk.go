package collector

import "syscall"

func DiskPercent() float64 {
	var stat syscall.Statfs_t
	if err := syscall.Statfs("/", &stat); err != nil || stat.Blocks == 0 {
		return 0
	}
	used := stat.Blocks - stat.Bfree
	return float64(used) / float64(stat.Blocks) * 100
}
