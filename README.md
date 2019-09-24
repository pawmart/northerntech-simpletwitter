# northern.tech Tweeter App

I have chosen go lang as it seemed the most robust to write it with. Swagger go was choosen to allow RAD development for the purpose of this demo.

## General comments and considerations

- Standard golang project structure as per [quidelines](https://github.com/golang-standards/project-layout)
- To follow [JSON API](https://jsonapi.org/) to speed things up I have used [go-swagger](https://github.com/go-swagger/go-swagger) auto-generation
- Storage is interfaced, so its easy to replace the storage with any type (and to allow ease of unit testing) as long as it meets the interface requirements (Ideally we should do the same for config and handler too, but for the purpose of the demo I'll omit it for now)
- Sorting direction filter for tweets list should be added too.
- Roles and permissions should be added for management endpoints not allowed for 'standard' user e.g. delete or patch...
- A user should be taken from the auth mechanisms rather then being passed as body
- Configuration should be likely fetched from secured place like 'Vault' or purely via env vars
- Storage (mongo) connection type and credentials are far from secure for production usage

## Instructions - Tests, runs, deployments

- We have local runs in Makefile as well as docker definitiions for mongo storage / app / go builder. This is to allow environment agnostic development and deployments.
- Unit tests and behavioural tests both executed using `go test ./...`

#### Run locally with:
- `go get github.com/pawmart/northerntech-simpletwitter`
- `cd ${GOPATH}/src/github.com/pawmart/northerntech-simpletwitter`
- `make up` (executing tests, builds and runs the service locally)

#### Deploy with:
- docker-compose.yml to docker swarm
- deployed to `66.70.186.117:6543/v1/tweets`

## Example

- Example get tweets request: GET: `http://localhost:6543/v1/tweets?filter[tag]=abc&count=2&year=2019`

## Misc

- Run `make` for the list of available commands.
- Scenarios for tests in `./api/features`
- Swagger definition in `/api/swagger` (can be imported to Postman to see exact detailed spec and test)
