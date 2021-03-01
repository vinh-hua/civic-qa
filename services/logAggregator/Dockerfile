# build stage
FROM golang:latest AS build-env
ADD . /src
RUN cd /src/cmd && go build -o goapp

# run stage
FROM  golang:latest
WORKDIR /app
COPY --from=build-env /src/cmd/goapp /app/
EXPOSE 8888
ENTRYPOINT ["./goapp"]