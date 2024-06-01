# Backend Template Go

This repository contains a Go backend project with a CRUD for clients.
What you can do here:

- Insert a large number of clients by .txt file;
- Insert, list, update and delete clients;
- List and search clients

### Prerequisites

- Make (optional)
- Docker (optional)
- Docker Compose (optional)
- Go 1.22.3 or latest

### Development Environment:

This repository mainly develop with docker and docker-compose environment for assuring the behaviour of the app itself.
Ps: I commited the .envs to make easyer for you.

You can run it locally by running `make docker.dev`

### Commands:

- `make docker.dev`: Build and run docker-compose development environment
- `make docker.dev.build`: Force build and run docker-compose development environment
- `make docker.build`: Build docker image
- `make docker.build.alpine`: Build docker image with alpine base
- `make clean`: Clean temporary and build folder
- `make build`: Build app

### Folder structure:

```
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── docker
│   ├── alpine.Dockerfile
│   ├── dev.Dockerfile
│   ├── docker-compose.dev.yaml
│   └── Dockerfile
├── Dockerfile -> ./docker/Dockerfile
├── go.mod
├── go.sum
├── go.work
├── internal
│   ├── crons
│   │   │── load.go
│   │   └── validate.go
│   ├── entities
│   │   └── client.go
│   ├── handlers
│   │   │── client.go
│   │   │── home.go
│   │   └── templates
│   │       └── index.html
│   ├── models
│   │   └── client.go
│   └── routes
│       └── routes.go
├── Makefile
├── pkg
│   ├── db
│   │   │── connection.go
│   │   └── migrations
│   └── utils
│       └── utils.go
├── README.md
└── tmp
    ├── build-errors.log
    └── main
```
