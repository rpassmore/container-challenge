# syntax=docker/dockerfile:1

########################################
# Build

FROM golang:1.17.2-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY main.go ./
COPY static/ static

RUN CGO_ENABLED=0 go build -o /main

#######################################
# Deploy
FROM scratch

WORKDIR /

COPY --from=build /main /main

EXPOSE 3000

#USER nonroot:nonroot

ENTRYPOINT ["/main"]

