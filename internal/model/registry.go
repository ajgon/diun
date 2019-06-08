package model

// Registry holds registry configuration
type Registry struct {
	Username    string `yaml:"username,omitempty" json:",omitempty"`
	Password    string `yaml:"password,omitempty" json:",omitempty"`
	InsecureTLS bool   `yaml:"insecure_tls,omitempty" json:",omitempty"`
	Timeout     int    `yaml:"timeout,omitempty" json:",omitempty"`
}