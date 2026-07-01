package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type NamedCheck struct {
	Name string `yaml:"name" json:"name"`
}
type Config struct {
	APIURL          string       `yaml:"api_url"`
	AssetID         string       `yaml:"asset_id"`
	Token           string       `yaml:"token"`
	IntervalSeconds int          `yaml:"interval_seconds"`
	VerifyTLS       bool         `yaml:"verify_tls"`
	ProcessChecks   []NamedCheck `yaml:"process_checks"`
	ServiceChecks   []NamedCheck `yaml:"service_checks"`
}

func Load(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	c := Config{IntervalSeconds: 60, VerifyTLS: true}
	section := ""
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(strings.Split(s.Text(), "#")[0])
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, ":") {
			section = strings.TrimSuffix(line, ":")
			continue
		}
		if strings.HasPrefix(line, "-") {
			name := strings.Trim(strings.TrimPrefix(line, "- name:"), " \"'")
			if section == "process_checks" {
				c.ProcessChecks = append(c.ProcessChecks, NamedCheck{Name: name})
			} else if section == "service_checks" {
				c.ServiceChecks = append(c.ServiceChecks, NamedCheck{Name: name})
			}
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key, val := parts[0], strings.Trim(strings.TrimSpace(parts[1]), "\"'")
		switch key {
		case "api_url":
			c.APIURL = val
		case "asset_id":
			c.AssetID = val
		case "token":
			c.Token = val
		case "interval_seconds":
			if n, err := strconv.Atoi(val); err == nil {
				c.IntervalSeconds = n
			}
		case "verify_tls":
			c.VerifyTLS = val != "false"
		}
	}
	return c, s.Err()
}
