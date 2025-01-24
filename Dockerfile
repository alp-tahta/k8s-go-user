# syntax=docker/dockerfile:1

FROM golang:1.23.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
# COPY go.mod go.sum ./ #Use this instead of above if you have third-party package
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /k8s-go-user

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8081

# Run
CMD ["/k8s-go-user"]