# dp-release-calendar-api

API for managing the release calendar

## Getting started

- Run `make debug`

### Dependencies

- No further dependencies other than those defined in `go.mod`

#### Tools

To run some of our tests you will need additional tooling:

##### Audit

We use `dis-vulncheck` to do auditing, which you will [need to install](https://github.com/ONSdigital/dis-vulncheck).

##### Linting

We use v2 of golangci-lint, which you will [need to install](https://golangci-lint.run/docs/welcome/install).

### Configuration

| Environment variable         | Default                     | Description                                                                                                        |
| ---------------------------- | --------------------------- | ------------------------------------------------------------------------------------------------------------------ |
| API_ROUTER_URL               | <http://localhost:23200/v1> | The URL of the [dp-api-router](https://github.com/ONSdigital/dp-api-router)                                        |
| BIND_ADDR                    | :27800                      | The host and port to bind to                                                                                       |
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s                          | The graceful shutdown timeout in seconds (`time.Duration` format)                                                  |
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s                         | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format) |
| HEALTHCHECK_INTERVAL         | 30s                         | Time between self-healthchecks (`time.Duration` format)                                                            |

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2023, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
