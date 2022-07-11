# linksy

Simple URL Shortener

### How to start

1. `make mod`
2. `make`

### Configuration

Edit `resources/config/config_default.yaml` or pass to `make` configuration file with `CONFIG_SOURCE_PATH` argument:

1. `service.db_type`: `memory` | `postgres`

### Tests

Acceptance tests can be run with `make acceptance`
