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

### Run locally

Create `.env` file in root directory and add following values:
```
PORT=
REDIS_HOST=
REDIS_PORT=
REDIS_PASS=
```

Use `docker-compose up` to spin up project

To see provided endpoints go to `https://localhost:<your_port>/swagger/index.html`

## TODO

- [X] Main logic: `/create_link`, `/redirect_by_short_link`, `/get_link_info`
- [X] config
- [X] error handling
- [ ] Tests: repository (redismock v9 not supported yet), config, integration tests
- [X] logging
- [X] linter
- [X] Swagger
- [X] Dockerize the app
- [X] Makefile
- [X] GitHub CI/CD
- [ ] README: local run steps, badges
- [X] LICENSE
- [ ] auth
- [ ] caching
- [ ] url validation
- [ ] non-http interaction?
- [ ] Custom logging?
- [ ] Stressful tests?
- [ ] telegram bot? (assigned to [@VSokolov2](https://github.com/VSokolov2))
- [ ] cli client?
- [ ] frontend?

## Contributing

Contributions are very welcome!

Feel free to open an issue or create a pull request with your additions.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
