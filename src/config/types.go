package config

type Profile struct {
	BaseUrl string            `yaml:"base_url"`
	Headers map[string]string `yaml:"headers"`
}
