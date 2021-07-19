# API

Placeholder for the API.  Right now this just runs on port 8080 and processes a few endpoints.

## Endpoints 

* `/activities/upload`: Uploads a workout via `POST`
* `/activities/{id}`: Retrieves a workout via `GET`

## Building

From the base directory, run 
`go generate ./...`
followed by 
`go build`