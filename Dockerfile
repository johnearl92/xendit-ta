##
## We build everything here
##
# FROM golang:alpine as build
FROM harbor.amihan.net/library/builder-go:latest as build

##
## Add git, ca-certificates and timezone info, needed if we call external services
##
# RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
# RUN apk update && apk add --no-cache tzdata 

##
## Add a new user here since we can't add it in scratch
##
ENV USER=or \
    UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

##
## Build the go binary here. CGO_ENABLED=0 to disable clib requirement for image to work in scratch
##
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
ARG VERSION=dev
WORKDIR /go/src/github.com/johnearl92/xendit-ta
COPY go.* ./
COPY . .
# RUN go build -mod vendor -o /go/bin/xendit-ta -ldflags "-X main.version=${VERSION} -w -s"
RUN GO111MODULE=on go build -mod=vendor -o /go/bin/xendit-ta -ldflags "-X main.version=${VERSION} -w -s"


##
## Final image uses scratch. We copy zoneinfo, ca-certs, user/group details, and the binary from the previous step
##
FROM scratch
# COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /go/bin/xendit-ta /xendit-ta
COPY swagger /swagger

##
## This image contains the migration files
##
# ADD db /db

##
## Set to the unprivileged user
##
USER or:or

##
## Set the binary as the entrypoint
##
ENTRYPOINT [ "/xendit-ta" ]
CMD [ "serve" ]