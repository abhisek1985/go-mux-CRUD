# Pull base image
FROM golang:1.17

# Create working directory
RUN mkdir /go/src/go-mux-CRUD

# Copy dependency file
COPY go.mod /go/src/go-mux-CRUD/
COPY go.sum /go/src/go-mux-CRUD/

# Set working directory for docker image
WORKDIR /go/src/go-mux-CRUD

# Download necessary Go modules
RUN go mod download

# Copy all files from current local directory to docker image working directory
COPY . /go/src/go-mux-CRUD

# Create executable file from our local main.go to docker image
RUN go build -o main /go/src/go-mux-CRUD

# Run excutable from docker image
ENTRYPOINT ["/go/src/go-mux-CRUD/main"]