//go:build linux
package service
func InstallHint() string { return "Create a systemd unit running: mortify-agent -config /etc/mortify-agent.yaml" }
