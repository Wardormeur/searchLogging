# Simplistic backend for logging

## Run
`go run server.go db.go`

## Test
### Integration
`go test ./...`
### Unit
TBD

## Entities
### Field
A named selector/value that have been scrapped

### Item
A representation of a selection with its n-fields, linked to a model

### Listing
Contains a list of Items to be fetched
