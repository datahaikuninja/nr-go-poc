# syntax=docker/dockerfile:1

ARG GO_VERSION

# Build the application from source
FROM golang:${GO_VERSION}-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /nrug-vol-09

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

ARG NR_APM_APP_NAME
ARG NR_LICENSE_KEY

ENV NR_APM_APP_NAME=${NR_APM_APP_NAME}
ENV NR_LICENSE_KEY=${NR_LICENSE_KEY}

WORKDIR /

COPY --from=build-stage /nrug-vol-09 /nrug-vol-09

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/nrug-vol-09"]
