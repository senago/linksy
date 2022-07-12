# linksy

Simple URL Shortener

### How to start

1. `make mod`
2. `make`

### Configuration

Edit [default config](resources/config/config_default.yaml) or pass to `make` configuration file with `CONFIG_SOURCE_PATH` argument

### Tests

[Acceptance tests](test/acceptance/shortener_test.go) can be run with `make acceptance`

### [Shortening Algorithm](internal/util/shortener.go)

1. Get the hash of an URL with some salt added (UUID is used), salt is used to produce different short links for the same input from potentially different users
2. Encode the hash into base58 and return first 10 characters
