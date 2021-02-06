# /app/mortar

Code in this directory is related to Mortar and uber-fx

- [Mortar Providers](https://github.com/go-masonry/mortar/tree/master/providers) invocations.
- Builders
- Your constructors wrapped as `fx.Option`
- Tracing/Monitoring/Logger implicit configurations

## HTTP

Constructors, options, interceptors and everything related to HTTP can be found in [http.go](http.go)

## Monitoring

This template is ready to work with a Prometheus service.
You need to configure it to expect a metric on this service default internal REST port.

```yml
mortar:
  # Web server related configuration
  server:
    ...
    rest:
      ...
      internal:
        port: 5382
```

By default it's `:5382/metrics`

Constructors, options and everything else can be found in [metrics.go](metrics.go).

## Tracing

This template is assuming you will use Jaeger service.

In order for the bundled client to connect your Jaeger service you need to export some ENVIRONMENT variables.

```sh
export JAEGER_AGENT_HOST = localhost
export JAEGER_AGENT_PORT = 6831
export JAEGER_SAMPLER_TYPE = const
export JAEGER_SAMPLER_PARAM = 1
```

They are included in the [`makefile`](../../Makefile), when you execute `make run` they are exported.

Constructors, options and everything else can be found in [tracing.go](tracing.go).

## Logger

This template is using the [zerolog](https://github.com/rs/zerolog) library.
Zerolog can be configured to output in JSON format which is useful in production or the Console format which is useful for debug/local environments.
By default it is configured to use the Console format.
You can export `SERVER_LOGGER_CONSOLE=false` environment variable to disable this behavior in production or change the [`config.yml`](../../config/config.yml) file.

Constructors, options and everything else can be found in [logger.go](logger.go).

## Config

This template is using the [viper](https://github.com/spf13/viper) library.

Constructors, options and everything else can be found in [config.go](config.go).

## Authentication

This template doesn't enforce any Authentication. However there is a sample code showing how you can parse a JWT if found.
Default usage can be found in [`../validations/auth.go`](../validations/auth.go)

Constructors, options and everything else can be found in [auth.go](auth.go).
