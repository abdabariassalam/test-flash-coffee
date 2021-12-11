package config

type (
	// Config .
	Config struct {
		// Base
		Base struct {
			Port string `yaml:"port"`
		} `yaml:"base"`

		// Postgresql.
		Postgresql struct {
			Url string `yaml:"url"`
		} `ymal:"postgresql"`

		// AWS S3
		AWS struct {
			Region          string `yaml:"region"`
			Bucket          string `yaml:"bucket"`
			AccessKeyID     string `yaml:"access_key_id"`
			SecretAccessKey string `yaml:"secret_access_key"`
		} `ymal:"aws"`
	}
)
