# godnd

[work in progress] go client for dnd api

## Development

While this is in development note the environment variable defined in `.envrc`
means we expect to be running everything against a local database.

## Notes

The `genny` package parses api documentation files from https://github.com/ecshreve/5e-srd-api/tree/master/src/views/partials
and generates `go` types based on the expected response objects.

## References

https://github.com/bagelbits/5e-srd-api
https://github.com/bagelbits/5e-database
