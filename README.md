# Go lang containerized 

Containarizing a go lang project using docker

my docker version is 24.0.4
Used docker image golang:1.20-alpine 

created a small project hellow binary world

Created docker container using docker-compose.yml

    version: '3'

    services:
        goapplcation:
            container_name: goapplcation
            build:
            context: ./app
            ports:
            - '8080:8080'




the docker file will be like

    FROM golang:1.20-alpine

    WORKDIR /app
    # in this project no extenal dependacy used so go.sum file not created so * used
    COPY go.mod go.sum* ./

    RUN go mod download

    WORKDIR $GOPATH/src/app

    COPY . .

    RUN go build -o main .

    CMD ["./main"]