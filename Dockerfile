FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go.mod download

COPY main.go ./

RUN go build -o /school-service-image

CMD [ "/school-service-image" ]
