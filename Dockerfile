FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN apt-get install gcc

COPY database/*.go ./database/
COPY handlers/*.go ./handlers/
COPY models/*.go ./models/
COPY data.db ./
COPY main.go ./


RUN CGO_ENABLED=1 GOOS=linux go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]