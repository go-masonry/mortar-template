package controllers

import (
	"context"

	helloworld "github.com/go-masonry/mortar-template/api"

	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type HelloWorldController interface {
	helloworld.GreeterServer
}

type helloworldControllerDeps struct {
	fx.In

	Logger log.Logger
}

type helloworldControllerImpl struct {
	*helloworld.UnimplementedGreeterServer // if keep this one added even when you change your interface this code will compile
	deps                                   helloworldControllerDeps
}

func CreateHelloworldController(deps helloworldControllerDeps) HelloWorldController {
	return &helloworldControllerImpl{
		deps: deps,
	}
}

func (w *helloworldControllerImpl) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	w.deps.Logger.Debug(ctx, "saying hello to %s", req.GetName())
	return &helloworld.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil
}
