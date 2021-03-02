# build stage: has everything needed to build an executable
FROM golang:latest AS build-env
ADD . /src
RUN cd /src/cmd/form_v0 && go build -o goapp

# run stage
FROM  golang:latest
WORKDIR /app
COPY --from=build-env /src/cmd/form_v0/goapp /app/
EXPOSE 7070
ENTRYPOINT ["./goapp"]