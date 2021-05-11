# Account

## Description
Handles Account creation, validation, login, logout, and session tracking.
## Features
Signup, Login, Logout, GetSession

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
- `SESS_IMPL`: `string`
    - Valid Session Store Implementation: "memory" | "redis"
- `REDIS_ADDR`: `string`
    - if `SESS_IMPL` is "redis", the redis connection address.
- `AGG_ADDR`: `string`
    - LogAggregator Service Address