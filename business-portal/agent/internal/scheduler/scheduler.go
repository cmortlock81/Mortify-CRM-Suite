package scheduler
import("time"; cfg "mortify-crm-agent/internal/config")
func Every(c cfg.Config, fn func()){for{fn(); time.Sleep(time.Duration(c.IntervalSeconds)*time.Second)}}
