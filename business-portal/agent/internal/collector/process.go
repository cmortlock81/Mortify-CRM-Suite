package collector

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	cfg "mortify-crm-agent/internal/config"
)

type CheckResult struct {
	Name    string `json:"name"`
	Running bool   `json:"running"`
}

func ProcessCount() int { entries, _ := os.ReadDir("/proc"); return len(entries) }
func Processes(checks []cfg.NamedCheck) []CheckResult {
	out := []CheckResult{}
	for _, c := range checks {
		out = append(out, CheckResult{Name: c.Name, Running: commandContains(c.Name)})
	}
	return out
}
func commandContains(name string) bool {
	cmd := "ps"
	args := []string{"-A", "-o", "comm="}
	if runtime.GOOS == "windows" {
		cmd = "tasklist"
		args = nil
	}
	out, err := exec.Command(cmd, args...).Output()
	return err == nil && strings.Contains(strings.ToLower(string(out)), strings.ToLower(name))
}
