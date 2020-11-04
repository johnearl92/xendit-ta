######################### START CONFIG ##########################
#
# INSTRUCTIONS: Change the appropriate values in this area.
# Don't change anything beyond this block if not deemed necessary
#
#################################################################
DOCKER_REGISTRY=
APP=xendit-ta
VERSION=latest
TAG=latest
PROJECT=jagbay01
IMAGE=${DOCKER_REGISTRY}${PROJECT}/${APP}:${TAG}

NAMESPACE=xendit-ta
HELM_CHART=xendit-ta
VALUES_FILE=helm/xendit-ta/values/dev.yaml

SWAGGER_HOST=swagger-local.json

########################## END CONFIG ###########################

all: clean lint build test package

clean: stop
	@echo '========================================'
	@echo ' Cleaning project'
	@echo '========================================'
	go clean
	rm -rf build | true
	@echo 'Done.'

deps:
	@echo '========================================'
	@echo ' Getting Dependencies'
	@echo '========================================'
	@echo 'Cleaning up dependency list...'
	go mod tidy
	@echo 'Downloading dependencies...'
	go mod download
	@echo 'Vendorizing project dependencies...'
	go mod vendor
	@echo 'Done.'

gen:
	@echo '========================================'
	@echo ' Generating dependencies'
	@echo '========================================'
	@go generate ./cmd
	@swagger generate spec -o ./swagger/swagger-generated.json
	@echo 'Done.'

generate-swagger:
	@echo '========================================'
	@echo ' Generating swagger '
	@echo '========================================'
	@swagger mixin ./swagger/swagger-generated.json ./swagger/${SWAGGER_HOST} -o ./swagger/swagger.json
	@echo 'Done.'

build: deps wire gen generate-swagger
	@echo '========================================'
	@echo ' Building project'
	@echo '========================================'
	go fmt ./...
	@go build -mod=vendor -o build/bin/${APP} -ldflags "-X main.version=${VERSION} -w -s" .
	@echo 'Done.'

test:
	@echo '========================================'
	@echo ' Running tests'
	@echo '========================================'
	@go test -mod=vendor ./...
	@echo 'Done.'

lint:
	@echo '========================================'
	@echo ' Running lint'
	@echo '========================================'
	@golint ./...
	@echo 'Done.'

run: build
	@echo '========================================'
	@echo ' Running application'
	@echo '========================================'
	@build/bin/${APP} serve
	@echo 'Done.'

package-chart:
	@echo '========================================'
	@echo ' Packaging chart'
	@echo '========================================'
	mkdir -p build/chart
	cp -r helm/${APP} build/chart
	helm package  --app-version ${APP} -u -d build/chart build/chart/${APP}
	@echo 'Done.'

package-image:
	@echo '========================================'
	@echo ' Packaging docker image'
	@echo '========================================'
	@docker build -t ${IMAGE} .
	@echo 'Done.'

package: package-image # package-chart

publish-chart: package-chart
	@echo '========================================'
	@echo ' Publishing chart'
	@echo '========================================'
	helm push build/chart/${APP} amihan
	@echo 'Done.'

publish: publish-image # publish-chart

start: deps gen generate-swagger
	@echo '========================================'
	@echo ' Starting application'
	@echo '========================================'
	docker-compose up -d
	@echo 'Done.'

stop:
	@echo '========================================'
	@echo ' Stopping application'
	@echo '========================================'
	docker-compose down
	@echo 'Done.'

tools:
	@echo 'installing tools...'
	go get github.com/google/wire/cmd/wire
	brew tap go-swagger/go-swagger
	brew install go-swagger

wire:
	wire gen cmd/wire.go

# DEVOPS
devops-harbor-login:
	@echo '========================================'
	@echo ' Harbor Login'
	@echo '${DOCKER_REGISTRY}'
	@echo '========================================'
	@echo ${HARBOR_PASS} | docker login ${DOCKER_REGISTRY} --username ${HARBOR_USER} --password-stdin
	@echo 'Done.'

devops-pull-image:
	@echo '========================================'
	@echo ' Getting latest image'
	@echo '========================================'
	docker pull ${IMAGE} || true

devops-package-image: generate-swagger
	@echo '========================================'
	@echo ' Packaging docker image'
	@echo '========================================'
	docker build -t ${IMAGE} .
	@echo 'Done.'
	
devops-publish-image:
	@echo '========================================'
	@echo ' Publishing image'
	@echo '========================================'
	docker push ${IMAGE}
	@echo 'Done.'

devops-setup-helm: 
	helm init --client-only
	helm repo add ${PROJECT} ${CHART_REPO} --username ${HARBOR_USER} --password ${HARBOR_PASS} 
	helm repo update
	@echo 'Done.'

deploy-helm:
	helm upgrade ${HELM_CHART} --install \
		--namespace ${NAMESPACE} \
		--set image.tag=${TAG} \
		--set extraLabels.git_hash=\"${CI_COMMIT_SHORT_SHA}\" \
		--values ${VALUES_FILE} helm/xendit-ta
	@echo 'Done.'

devops-deploy-chart-manual:
	helm upgrade ${HELM_CHART} --install \
		--namespace ${NAMESPACE} \
		--set registries[0].url=${DOCKER_REGISTRY} \
		--set registries[0].username=amihan-robot \
		--set registries[0].password=AmihanGlobal101 \
		--set extraLabels.git_hash=\"00000001\" \
		--values ${VALUES_FILE} helm/xendit-ta
	@echo 'Done.'