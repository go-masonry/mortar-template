package validations

import (
	"context"
	"fmt"

	helloworld "github.com/go-masonry/mortar-template/api"
)

type HelloworldValidations interface {
	helloworld.GreeterServer
}

type helloworldValidationsImpl struct {
	*helloworld.UnimplementedGreeterServer
}

func CreateHelloworldValidations() HelloworldValidations {
	return new(helloworldValidationsImpl)
}

func (impl *helloworldValidationsImpl) SayHello(ctx context.Context, req *helloworld.HelloRequest) (_ *helloworld.HelloReply, err error) {
	if len(req.GetName()) == 0 {
		return nil, fmt.Errorf("name cannot be empty")
	}
	return nil, nil
}
