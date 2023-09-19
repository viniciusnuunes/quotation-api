FROM golang:1.21

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD [ "air", "-c", ".air.toml" ]
# WORKDIR /app/cmd/quotation-api

# RUN go build -o main .

# EXPOSE 8080

# CMD [ "./main" ]

# ENTRYPOINT [ "air" ]