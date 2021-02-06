package mortar

import (
	"github.com/go-masonry/mortar/constructors"
	"go.uber.org/fx"
)

func AuthFxOptions() fx.Option {
	return fx.Options(
		fx.Provide(constructors.DefaultJWTTokenExtractor),
	)
}
