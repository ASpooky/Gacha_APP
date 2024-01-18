FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest

#CMD ["air"]