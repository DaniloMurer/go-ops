# go-ops

## Setup

install go for your os

## Run

run the go file with following parameters:
```bash
go run main.go YOUR_DEPLOYMENT_FILE OLD_IMAGE_VERSION NEW_IMAGE_VERSION
```
example:
```bash
go run main.go ./deployment.yaml churrer.xyz:0.5.0 churrer.xyz:0.6.0
```
Note that the current implementation pretty much replaces anything with everything as long as your parameters are valid (content exists in file and filepath is valid)
