package validations

import (
	"context"
	"fmt"

	helloworld "github.com/go-masonry/mortar-template/api"
	"github.com/go-masonry/mortar/interfaces/auth/jwt"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type HelloworldValidations interface {
	helloworld.GreeterServer
}

type helloworldValidationsDeps struct {
	fx.In

	jwt.TokenExtractor
	Logger log.Logger
}

type helloworldValidationsImpl struct {
	*helloworld.UnimplementedGreeterServer
	deps helloworldValidationsDeps
}

func CreateHelloworldValidations(deps helloworldValidationsDeps) HelloworldValidations {
	return &helloworldValidationsImpl{
		deps: deps,
	}
}

func (impl *helloworldValidationsImpl) SayHello(ctx context.Context, req *helloworld.HelloRequest) (_ *helloworld.HelloReply, err error) {
	if err = impl.CheckAuth(ctx); err != nil {
		return nil, err
	}
	if len(req.GetName()) == 0 {
		return nil, fmt.Errorf("name cannot be empty")
	}
	return nil, nil
}
