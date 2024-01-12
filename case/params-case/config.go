package params_case

type Config struct {
	Name string `json:"name"`
}

func NewConfig() *Config {
	return &Config{
		Name: "params-case",
	}
}
