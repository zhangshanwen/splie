package conf

type Distributed struct {
	Setting Setting `yaml:"setting"`
}
type Setting struct {
	Name         string `yaml:"name"`
	PingSwitch   string `yaml:"ping_switch"`
	PingDuration int    `yaml:"ping_duration"`
}
