FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/quotation-api

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]