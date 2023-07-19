FROM golang:1.20-alpine
WORKDIR /app
# in this project no extenal dependacy used so go.sum file not created so * used
COPY go.mod go.sum* ./

RUN go mod download

WORKDIR $GOPATH/src/app

COPY . .

RUN go build -o main .

CMD ["./main"]