package collector
import("strings";"github.com/shirou/gopsutil/v3/process"; cfg "mortify-crm-agent/internal/config")
type CheckResult struct{Name string `json:"name"`; Running bool `json:"running"`}
func ProcessCount()int{p,_:=process.Processes(); return len(p)}
func Processes(checks []cfg.NamedCheck)[]CheckResult{ps,_:=process.Processes(); names:=map[string]bool{}; for _,p:=range ps{n,_:=p.Name(); names[strings.ToLower(n)]=true}; out:=[]CheckResult{}; for _,c:=range checks{out=append(out,CheckResult{c.Name,names[strings.ToLower(c.Name)]})}; return out}
