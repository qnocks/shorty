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
- ...

Create `.env` file in root directory and add following values:
```
PORT=
REDIS_HOST=
REDIS_PORT=
REDIS_PASS=
```

Use `make run` to build & run project, `make lint` to check code with linter.

## TODO

- [X] Main logic: `/create_link`, `/redirect_by_short_link`, `/get_link_info`
- [X] config
- [X] error handling
- [ ] Tests: Service (link), repository (link), config
- [X] logging
- [X] linter
- [X] Open API
- [ ] Dockerize the app
- [ ] Stressful tests
- [ ] Makefile
- [ ] GitHub CI/CD
- [ ] Custom logging
- [ ] auth?
- [ ] cli client?
- [ ] README
- [ ] LICENSE

## License

> Coming soon
