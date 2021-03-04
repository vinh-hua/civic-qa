# build stage: has everything needed to build an executable
FROM golang:latest AS build-env
ARG PORT=7070
ADD . /src
RUN cd /src/cmd/form_v0 && go build -o goapp

# run stage
FROM  golang:latest
ARG PORT=7070
WORKDIR /app
COPY --from=build-env /src/cmd/form_v0/goapp /app/
COPY --from=build-env /src/templates /templates/
EXPOSE ${PORT}
ENTRYPOINT ["./goapp"]