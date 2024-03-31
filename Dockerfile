# syntax=docker/dockerfile:1

################################
#          Build step          #
################################
FROM golang:1.22-alpine as build

# Set destination for copy commands down below
WORKDIR /app

# Copy the go metafiles & download all packages
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /duckploy && \
    chmod +x /duckploy


################################
#         Release step         #
################################
FROM gcr.io/distroless/static-debian12 as release

# Set destination for copy commands down below
WORKDIR /

# Copy build binary into release stage
COPY --from=build /duckploy /duckploy

# Specify nonroot user
USER nonroot:nonroot

# Specify command
ENTRYPOINT ["/duckploy"]
