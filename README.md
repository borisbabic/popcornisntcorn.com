# popcornisntcorn.com

The code (`main.go`) is too simplistic to be useful the rest is probably over-engineered.

This is also improper image handling, however the needs are very simple so `¯\_(ツ)_/¯`

## Build

### Golang
To create a static go binary build with 
`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o notcorn` 

You can now run `./notcorn`

### Docker

After building the go binary build 
`docker build . -t cloud.canister.io:5000/babic/popcornisntcorn:latest`

Then run `docker-compose up` to get it up.

# FAQ
## What's with the labels in the docker-compose.yml
I'm deploying this to my personal rancher setup that uses traefik. That's what the labels are for and why the file is located in `deploy/notcorn`
## What's with the docker-compose.override.yml
It's done like so to isolate the stuff needed locally, we don't want the port open in "production" but we do locally. 
