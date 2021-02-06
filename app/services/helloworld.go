package services

import (
	"context"

	helloworld "github.com/go-masonry/mortar-template/api"
	"github.com/go-masonry/mortar-template/app/controllers"
	"github.com/go-masonry/mortar-template/app/validations"
	"github.com/go-masonry/mortar/interfaces/monitor"

	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type helloworldServiceDeps struct {
	fx.In

	Logger log.Logger

	Validations validations.HelloworldValidations
	Controller  controllers.HelloWorldController
	Metrics     monitor.Metrics `optional:"true"`
}

type helloworldServiceImpl struct {
	helloworld.UnimplementedGreeterServer // if keep this one added even when you change your interface this code will compile
	deps                                  helloworldServiceDeps
}

func CreateHelloworldService(deps helloworldServiceDeps) helloworld.GreeterServer {
	return &helloworldServiceImpl{
		deps: deps,
	}
}

func (impl *helloworldServiceImpl) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	_, err := impl.deps.Validations.SayHello(ctx, req)
	if err != nil {
		impl.deps.Logger.WithError(err).Debug(ctx, "validation failed")
		return nil, err
	}
	return impl.deps.Controller.SayHello(ctx, req)
}
