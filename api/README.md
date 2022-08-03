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

## New notes
docker build -f ./Dockerfile -t run_log_api .
docker run -it -p 8080:8080 run_log_api

docker-compose run --service-ports --entrypoint bash run-log-api

docker-compose run --service-ports --entrypoint bash run-log-api
curl -X POST --data-binary @/home/emurray/Downloads/Morning_Run.fit http://localhost:8080/activity/upload
curl -X POST -d 'name=testname' http://localhost:8080/user/create
curl -i http://localhost:8080/api/v1/users/1/activities/1