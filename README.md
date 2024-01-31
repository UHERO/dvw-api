# DVW-API
A REST API built with Go to support DBEDT's data warehouse.

# Running & Building Locally
Clone the repo to your machine.

Run the following in the project's root dir to initilize or update the go mod file
```
go mod init github.com/UHERO/dvw-api
go mod tidy
```

To run a Go app locally, use
```
go run main.go
```

However, as far as I'm aware, the DVW-API isn't configured for local development. It will fail to connect to the prod database. I would like to setup a test environment at some point to make development feasible.

run the following command to build the executable for running on our linux box
```
env GOOS="linux" GOARCH="amd64" go build
```

# ReadMe ToDos
[ ] Add Diagram Illustration for how dvw-api relates to DBEDT and UHERO services
[ ] Add Deployment Process
[ ] Configure a way to run locally
[ ] Expand About section to elaborate on the purpose of the api