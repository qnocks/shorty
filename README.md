# Shorty

## Overview

Application allows to reduce long links and track statistics for your business and projects by
monitoring the number of hits from your URL with the click counter

### Build with

- Go 1.18
- gin-gonic/gin
- Redis
- Docker, docker-compose

## Usage

### Prerequisites

- go 1.18
- docker & docker-compose
- [golangci-lint](https://github.com/golangci/golangci-lint) (<i>optional</i>, used to run code checks)
- [swag](https://github.com/swaggo/swag) (<i>optional</i>, used to re-generate swagger documentation)

- ...

Create `.env` file in root directory and add following values:
```
PORT=
REDIS_HOST=<same_as_redis_container_name>
REDIS_PORT=
REDIS_PASS=
```

Use `make run` to build & run project, `make lint` to check code with linter.

## TODO

- [X] Main logic: `/create_link`, `/redirect_by_short_link`, `/get_link_info`
- [X] config
- [X] error handling
- [ ] Tests: repository - redismock v9 not supported yet, config
- [X] logging
- [X] linter
- [X] Open API
- [X] Dockerize the app
- [ ] Stressful tests
- [X] Makefile
- [ ] GitHub CI/CD
- [ ] Custom logging
- [ ] auth?
- [ ] cli client?
- [ ] README
- [ ] LICENSE

## License

> Coming soon
