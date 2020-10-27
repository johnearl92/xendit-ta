XENDIT Technical Assessment
--------------------

This is a project implementation of the technical assessment 


## Prerequisire

### Workspace

- Go
- Docker
- Docker hub
- Make

### Docker images

- jagbay01/xendit-ta

### Deployment

- Docker Compose
- Kubernetes
- Minikube
- Helm

## Build
*If wire tools is not yet installed execute `make tools`*

Run the following to build the binary:

```shell
make build
```

The binary will be located in: `build/bin/xendit-ta`

## Configuration

Configuration is handled by [viper](https://github.com/spf13/viper) which allows configuration using a config file or by environment variables.

A sample configuration file can be found in `config/.xendit-ta.yaml`. Copy this to the `$HOME` directory to override the defaults.

Setting environment variables can also override the default configuration:

| Environment Variable                 | Description                                      | Default                   |
| ------------------------------------ | ------------------------------------------------ | ------------------------- |
| XENDIT_TA_SERVICE_SERVER_HOST | Name of the host or interface to bind the server | 0.0.0.0                   |
| XENDIT_TA_SERVICE_SERVER_PORT | Port to bind the server                          | 8080                      |

## Starting the Service

```shell
make run
```

## Packaging Image

To create the docker image:

```
make package
```

To publish the image to docker hub:

```
make publish
```

Note that publishing the image requires access to jagbay01 account in docker hub.


### Project Structure

```
xendit-ta
|- build/               # build artifacts are generated here
|- cmd/                 # command line commands live here. Checkout cobra library
|- config/              # configuration files are here
|- db/                  # for database migration files
|- helm/                # helm chart for kubernetes deployment
|- internal/            # for internal go packages 
| |- server
| |- ...
|- swagger              # swagger ui
| |- swagger.json       # swagger spec generated
|- .dockerignore        # ignore list for docker
|- .gitignore           # ignore list for git
|- go.mod               # dependencies for project
|- go.sum               # checksum for dependencies, do not manually change
|- main.go              # the main go file
|- Makefile             # build scripts
|- README.md            # this file
```

### Adding Dependencies

To add dependencies, run the following:

```shell
go get -u {dependency}
make deps
``` 


### Docker Compose Deployment
```shell
make package start
```

package - will create the docker image
start - will start the docker-compose 