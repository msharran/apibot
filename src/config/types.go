package config

type Profile struct {
	BaseUrl  string             `yaml:"base_url"`
	Headers  map[string]string  `yaml:"headers"`
	Requests map[string]Request `yaml:"requests"`
}

type Request struct {
	Endpoint string `yaml:"endpoint"`
	Method   string `yaml:"method"`
	Body     string `yaml:"body"`
}
