FROM golang:1.19-alpine as builder

# Create and change to the app directory.
WORKDIR /webfinger

# Expecting to copy go.mod and if present go.sum.
COPY *.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app ./cmd/main.go

# Install dependencies on the target image
FROM alpine
RUN apk --no-cache add ca-certificates

WORKDIR /webfinger

# Copy the binary to the production image from the builder stage.
COPY --from=builder /webfinger/app /webfinger

# Copy the payload to the production image from builder
COPY --from=builder /webfinger/data/payload.json /webfinger/data/

# Intended entryport 
EXPOSE 8080

# Create a non-root user and run the application in that userspace
RUN addgroup -S apprunner
RUN adduser --disabled-password --gecos "" --ingroup apprunner apprunner
USER apprunner 

# Run the web service on container startup.
CMD ["./app"]