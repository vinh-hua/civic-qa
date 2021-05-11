# gateway

## Description
Primary API Gateway/Facade for subservices
## Features
Routes Requests to appropriate subservice, adds authentication information to requests.

## API
OpenAPI: `api-v0.yml`

## Variables:
env:
- `ADDR`: `string`
    - Service address/port
- `ACCOUNT_SVC`: `string`
    - Account service address/port
- `FORM_SVC`: `string`
    - Form service address/port
- `MAILTO_SVC`: `string`
    - Mailto service address/port
- `AGG_ADDR`: `string`
    - LogAggregator Service Address