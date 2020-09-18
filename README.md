# [dynamicdns](https://git.ryanburnette.com/ryanburnette/dynamicdns)

A dynamic DNS service written in Go.

## Configuration

Configure by creating a `.env` file next to the binary.

### General

```
# .env
DOMAIN=[fqdn]
SCHEDULE=[cron expression]
```

### Name.com

```
# .env
API=NAMECOM
NAMECOM_USERNAME=[username]
NAMECOM_API_TOKEN=[apitoken]
```

## Implementation

Make it a service with [go-serviceman](https://git.coolaj86.com/coolaj86/go-serviceman).
