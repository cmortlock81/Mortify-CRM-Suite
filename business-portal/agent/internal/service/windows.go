//go:build windows
package service
func InstallHint() string { return "Install with sc.exe create MortifyAgent binPath= <path>" }
