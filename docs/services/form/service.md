# form

## Description
Handles the creation and management of forms, responses, and tags
## Features
Forms, Responses, Tags

## API
OpenAPI: `api-v0.yml`

## Variables:
env:
- `ADDR`: `string`
    - Service address/port
- `DB_IMPL`: `string`
    - Valid DB Implementation: "sqlite" | "mysql"
- `DB_DSN`: `string`
    - Primary DB DSN
- `ANALYTICS_IMPL`: `string`
    - Valid Analytics Service Client Implementation: "v0"
- `AGG_ADDR`: `string`
    - LogAggregator Service Address