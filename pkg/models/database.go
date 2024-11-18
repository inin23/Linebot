package models

type Config struct {
	Agent struct {
		Authtoken string `yaml:"authtoken"`
	} `yaml:"agent"`

	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
	} `yaml:"database"`
}
