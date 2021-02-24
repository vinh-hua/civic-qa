# build stage
FROM golang:latest AS build-env
ADD . /src
RUN cd /src/cmd && go build -o goapp

# run stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/cmd/goapp /app/
ENTRYPOINT ["./goapp"]