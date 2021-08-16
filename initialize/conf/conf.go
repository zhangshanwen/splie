package conf

type Conf struct {
	Port          string        `yaml:"port"`
	DB            DB            `yaml:"db"`
	Authorization Authorization `yaml:"authorization"`
	Distributed   Distributed   `yaml:"distributed"`
	Level         string        `yaml:"level"`
}
