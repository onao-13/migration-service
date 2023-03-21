# STAGE 1: COPY APP AND DOWNLOAD DEPENDENCIES
FROM golang:1.19.4-alpine3.17 AS build

WORKDIR /service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY interanls ./internals

# STAGE 2: BUILD IMAGE
FROM scratch

COPY --from=build /service/migration ./migration
EXPOSE 8083
ENTRYPOINT ["/migration"]

# STAGE 3: COPY STATIC FILES
COPY sql ./sql