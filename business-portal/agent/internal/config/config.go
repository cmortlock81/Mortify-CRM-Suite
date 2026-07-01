package config
import("os";"gopkg.in/yaml.v3")
type NamedCheck struct{Name string `yaml:"name" json:"name"`}
type Config struct{APIURL string `yaml:"api_url"`; AssetID string `yaml:"asset_id"`; Token string `yaml:"token"`; IntervalSeconds int `yaml:"interval_seconds"`; VerifyTLS bool `yaml:"verify_tls"`; ProcessChecks []NamedCheck `yaml:"process_checks"`; ServiceChecks []NamedCheck `yaml:"service_checks"`}
func Load(path string)(Config,error){b,e:=os.ReadFile(path); if e!=nil{return Config{},e}; var c Config; e=yaml.Unmarshal(b,&c); if c.IntervalSeconds==0{c.IntervalSeconds=60}; return c,e}
