package params_case

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type ServiceB struct {
	Name string `json:"name"`
}

func NewServiceB(lc fx.Lifecycle, config *Config) *ServiceB {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("ServiceB OnStart...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("ServiceB OnStop...")
			return nil
		},
	})
	return &ServiceB{
		Name: config.Name,
	}
}

func (s *ServiceB) DoSomething() {
	fmt.Println(s.Name)
	fmt.Println("DoSomething...")
}

func (s *ServiceB) ExecuteTask() {
	// 执行任务的逻辑
	fmt.Println("ExecuteTask...")
}
