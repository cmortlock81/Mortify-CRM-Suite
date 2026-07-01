package collector

import (
	"os/exec"
	"runtime"
	"strings"

	cfg "mortify-crm-agent/internal/config"
)

func Services(checks []cfg.NamedCheck) []CheckResult {
	out := []CheckResult{}
	for _, c := range checks {
		out = append(out, CheckResult{Name: c.Name, Running: serviceRunning(c.Name)})
	}
	return out
}

func serviceRunning(name string) bool {
	if runtime.GOOS == "linux" {
		if err := exec.Command("systemctl", "is-active", "--quiet", name).Run(); err == nil {
			return true
		}
	}
	if runtime.GOOS == "windows" {
		out, err := exec.Command("sc", "query", name).CombinedOutput()
		return err == nil && strings.Contains(strings.ToLower(string(out)), "running")
	}
	return commandContains(name)
}
