# logAggregator

## Description:
This service is used to aggregate logs in a common format from many sources.

## Features:
Log: store a new log, Query: query existing logs

## API:
OpenAPI: `api-v0.yaml`

## Variables
env:
- `ADDR`: `string`
    - server addr/port
- `DN_DSN`: `string`
    - LOG DB DSN (not main db)