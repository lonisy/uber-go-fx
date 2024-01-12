package params_case

import (
	"context"
	"go.uber.org/fx"
	"log"
)

func Run2() {
	//fx.New(fx.Provide(NewConfig, NewServiceA, NewServiceB), fx.Invoke(func(s *ServiceB) {
	//	s.DoSomething()
	//})).Run()

	// Todo 这个示例代码，会执行结束，适合一些命令或者任务的脚本程序
	app := fx.New(
		// 提供服务和组件
		fx.Provide(NewConfig, NewServiceA, NewServiceB),
		// 调用启动逻辑
		fx.Invoke(
			func(s *ServiceB) {
				s.ExecuteTask()
			},
		),
	)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := app.Stop(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// 检查启动是否有错误
	if err := app.Err(); err != nil {
		log.Fatal(err)
	}
}

func Run() {
	//fx.New(fx.Provide(NewConfig, NewServiceA, NewServiceB), fx.Invoke(func(s *ServiceB) {
	//	s.DoSomething()
	//})).Run()

	// 这段代码不会自动结束，需要手动结束
	app := fx.New(
		// 提供服务和组件
		fx.Provide(NewConfig, NewServiceA, NewServiceB),
		// 调用启动逻辑
		fx.Invoke(
			func(s *ServiceB) {
				s.ExecuteTask()
			},
		),
	)
	app.Run()
}
