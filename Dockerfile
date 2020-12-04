FROM golang:1.14.12
WORKDIR /omdbapp
COPY . /omdbapp

# Install dependencies
RUN go mod vendor

# Compile
RUN go build -o main .

# Run
ENTRYPOINT "/omdbapp/main"
