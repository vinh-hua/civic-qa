# build stage
FROM golang:latest AS build-env
ARG PORT=8888
ADD . /src
RUN cd /src/cmd && go build -o goapp

# run stage
FROM  golang:latest
ARG PORT=8888
WORKDIR /app
COPY --from=build-env /src/cmd/goapp /app/
EXPOSE ${PORT}
ENTRYPOINT ["./goapp"]