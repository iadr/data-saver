FROM golang:1.22.5 as builder

WORKDIR /project

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /project/values-importer ./cmd/main

ENTRYPOINT [ "/project/values-importer" ]
