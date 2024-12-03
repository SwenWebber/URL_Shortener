#official image Go 1.23.1 
FROM golang:1.23.1-alpine

#installing required tools
RUN apk add --no-cache git make

#creating workdir
WORKDIR /app
#copying dependency files
COPY go.mod go.sum ./


#downloading dependencies
RUN go mod download


#copy source code
COPY . .
#exposing port
EXPOSE 8080
#running the app
CMD ["go","run", "cmd/main.go"]