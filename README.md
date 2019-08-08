# Dockerized Development Environment and Tools for full-stack gRPC service generation.

[![Build Status](https://travis-ci.com/magnuskokk/grpc-tools.svg?branch=master)](https://travis-ci.com/magnuskokk/grpc-tools)

This project aims to set up an opinionated environment for protobuf based web development, mainly to generate and prototype remote monitoring and control panel systems for abstract devices.

## Ideas
This project is inspired by https://github.com/gogo/grpc-example and uses https://github.com/uber/prototool to lint and generate protos.

Imagine you have a custom home automation device (raspi with a temperature sensor and a radiator). It may have its own internal logic to turn on/off the heat based on indoor temperature but that's not important here. We need to remotely see if everything's working and intervene if needed.

The next idea would be to research how to generate Prometheus and Grafana provisioning scripts using a custom protoc plugin and custom annotations.

Example proto proto based on a real service from `idl/raspi/raspiv1`
```protobuf
message Temperature {
  sint32 reading = 1;
}

message Radiator {
  bool enabled = 1;
  uint32 level = 2;
}

message Status {
  Temperature temperature = 1;
  Radiator radiator = 2;
}
```

## Required packages for development
* `direnv`
* `docker`
* `docker-compose`

## Setup the development environment
The environment is currently simply a task runner and a golang docker image with local user.
* `$ git clone https://github.com/magnuskokk/grpc-tools.git`
* `$ cd grpc-tools`
* Install `direnv` from your package manager and set up the shell hook for the terminal emulator you're using (bash, zsh, etc...): https://github.com/direnv/direnv.
* `$ direnv allow .` to load local environment variables from .envrc.
* `$ ./install-tusk.sh` to install the task runner.
* `$ tusk` to list all tasks.
* `$ tusk env.build` to build the development environment images.

## Project layout
* All proto services are defined in `./idl/{name}/{name}{version}`. The generated go package for each service is `app/idl/{name}/{name}{version}`.
* Implementations of services are in `./app/api/{name}`.
* Server main files and Dockerfiles are generated into `./app/cmd/{name}-{type}-server` where type is either `grpc` or `gateway`.
* Swagger docs are generated into `./swagger/idl/{name}/{name}{version}`.
* Typescript client is generated into `./frontend/generated` (this needs work).

Creating new services is just a matter of creating a `.proto` file, running generators, implementing the package under `api` and creating an entry in `docker-compose.yml` so it can run.

There are two example services `raspi` and `echo`.

## List all tasks
* `$ tusk`
```
Tasks:
   docker.clean             Stop and clean everything but keep system cache for fast rebuilds.
   docker.clean.containers  Stop and remove all containers, images and any anonymous volumes attached to containers.
   docker.clean.images      Stop all containers and remove all images.
   docker.clean.volumes     Stop all containers and remove all volumes.
   docker.down              Stop all containers. All docker.* commands include only services defined in docker-compose.yml file.
   docker.prune.system      WARNING! Prunes the whole system.
   env.build                Build containers for dev tools.
   env.clean                Reset any environment in this directory.
   gen.all                  Run all generators. (Must have run gen.install.tools first).
   gen.clean                Remove ALL generated files and directories.
   gen.protoc               Generate gRPC server, client, gateway, typescript and swagger for all services.
   gen.protoc.clean         Clean all protoc generated files.
   gen.protolint            Lint all protobuf definitions using prototool.
   gen.servers              Parse proto API services from ./idl and generate server main files and Dockerfiles.
   gen.servers.clean        Clean generated servers from ./app/cmd
   go.bench                 Run go benchmarks.
   go.generate              Generate a go package using //go:generate.
   go.generate.clean        Clean up files generated by //go:generate.
   go.install.tools         Install tools and dependencies necessary for code generation.
   go.test                  Run go tests. For example go.test ./...
   go.vet                   Run go vet
   stack.build              Build the whole stack.
 ```

## Build and run the stack
* `$ tusk stack.build`
* `$ docker-compose up -d`
* Make sure it's responding: `$ curl http://localhost:8000/echo?message=test` 

## Swagger UI with "Try out"
* echo server: `http://localhost:8080`
* raspi server: `http://localhost:8081`

## Reset the dev environment:
* `$ tusk env.clean.sudo`
* `$ tusk env.build`

### VSCode
direnv extension: https://marketplace.visualstudio.com/items?itemName=Rubymaniac.vscode-direnv. This allows installing all go tools in the .direnv directory.
