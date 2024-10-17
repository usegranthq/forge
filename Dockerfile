# Set the arg variables
ARG PORT

#### Base stage

# use the official Golang image as a build stage
FROM golang:1.23-alpine AS base

# install necessary packages
RUN apk add --no-cache make

# set the current working directory inside the container
WORKDIR /app

# copy go mod and sum files
COPY go.mod go.sum ./

# download all dependencies.
# dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

### Development stage
FROM base AS development

# install Air
RUN go install github.com/air-verse/air@latest

# copy the source from the current directory to the working directory inside the container
COPY . .

# create a tmp directory
RUN mkdir -p /app/tmp

# expose port to the outside world
EXPOSE $PORT

# command to run Air and hot reload changes
CMD ["air", "-c", ".air.toml"]

### Build stage
FROM base AS build

# copy the source from the current directory to the working directory inside the container
COPY . .

# build the binary
RUN go build -o /app/main cmd/api/main.go

### Production stage
FROM scratch AS production

# set the current working directory inside the container
WORKDIR /app

# copy the binary from the build stage
COPY --from=build /app/main /app/main

# expose port to the outside world
EXPOSE $PORT

# run the server
CMD ["/app/main"]
