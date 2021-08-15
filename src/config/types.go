package config

type Profile struct {
	BaseUrl             string `yaml:"base_url"`
	AuthorizationHeader string `yaml:"authorization_header"`
}
