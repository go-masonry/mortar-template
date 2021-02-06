package mortar

import (
	"os"

	"github.com/go-masonry/bjaeger"
	"github.com/go-masonry/bzerolog"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/providers"
	"go.uber.org/fx"
)

func LoggerFxOption() fx.Option {
	return fx.Options(
		fx.Provide(zeroLogBuilder),
		providers.LoggerFxOption(),
		providers.LoggerGRPCIncomingContextExtractorFxOption(),
		bjaeger.TraceInfoContextExtractorFxOption(),
	)
}

func zeroLogBuilder(config cfg.Config) log.Builder {
	builder := bzerolog.Builder().IncludeCaller()
	if config.Get("server.logger.console").Bool() {
		builder = builder.
			SetWriter(bzerolog.ConsoleWriter(os.Stderr))
	}
	return builder
}
