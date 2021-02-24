# build stage: has everything needed to build an executable
FROM golang:latest AS build-env
ADD . /src
RUN cd /src/cmd/auth_v0 && go build -o goapp

# run stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/cmd/auth_v0/goapp /app/
ENTRYPOINT ["./goapp"]