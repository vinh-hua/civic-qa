# build stage: has everything needed to build an executable
FROM golang:latest AS build-env
ARG PORT=8080
ADD . /src
RUN cd /src/cmd/account_v0 && go build -o goapp

# run stage
FROM  golang:latest
ARG PORT=8080
WORKDIR /app
COPY --from=build-env /src/cmd/account_v0/goapp /app/
EXPOSE ${PORT}
ENTRYPOINT ["./goapp"]