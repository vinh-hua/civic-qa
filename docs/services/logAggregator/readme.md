# logAggregator

## About:
This service is used to aggregate logs in a common format from many sources.

## Features:
- Log: store a new log
- Query: query existing logs

## API:
OpenAPI: `api-v0.yaml`
            
## Build:
navigate to `/dev/`, execute `$make build`

## Test:
navigate to `/dev/`, execute `$make test` or `$make test_verbose`

## Variables:
env:
- ADDR (default = :8888)
    - server port
- dbPath (default = logs.db)
    - path to sqlite database file