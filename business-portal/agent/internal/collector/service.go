package collector
import cfg "mortify-crm-agent/internal/config"
func Services(checks []cfg.NamedCheck)[]CheckResult{out:=[]CheckResult{}; for _,c:=range checks{out=append(out,CheckResult{Name:c.Name,Running:false})}; return out}
