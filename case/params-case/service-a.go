package params_case

type ServiceA struct {
	Name string `json:"name"`
}

func NewServiceA(config *Config) *ServiceA {
	return &ServiceA{
		Name: config.Name,
	}
}
